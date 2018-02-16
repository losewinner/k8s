/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pluginwatcher

import (
	"fmt"
	"net"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	registerapi "k8s.io/kubernetes/pkg/kubelet/apis/pluginregistration/v1alpha1"
	utilfs "k8s.io/kubernetes/pkg/util/filesystem"
)

// RegisterCallbackFn is the type of the callback function that handlers will provide
type RegisterCallbackFn func(pluginName string, versions []string, socketPath string) error

// Watcher is the plugin watcher
type Watcher struct {
	path     string
	handlers map[string]RegisterCallbackFn
	stopCh   chan interface{}
	fs       utilfs.Filesystem
	watcher  *fsnotify.Watcher
	wg       sync.WaitGroup
	mutex    sync.Mutex
}

// NewWatcher provides a new watcher
func NewWatcher(sockDir string) Watcher {
	return Watcher{
		path:     sockDir,
		handlers: make(map[string]RegisterCallbackFn),
		fs:       &utilfs.DefaultFs{},
	}
}

// AddHandler registers a callback to be invoked for a particular type of plugin
func (w *Watcher) AddHandler(handlerType string, handlerCbkFn RegisterCallbackFn) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.handlers[handlerType] = handlerCbkFn
}

// Creates the plugin directory, if it doesn't already exist.
func (w *Watcher) createPluginDir() error {
	glog.Infof("Ensuring Plugin directory at %s ", w.path)
	err := w.fs.MkdirAll(w.path, 0755)
	if err != nil {
		return fmt.Errorf("error (re-)creating driver directory: %s", err)
	}
	return nil
}

// Walks through the plugin directory to discover any existing plugin sockets.
func (w *Watcher) traversePluginDir() error {
	files, err := w.fs.ReadDir(w.path)
	if err != nil {
		return fmt.Errorf("error reading the plugin directory: %v", err)
	}
	for _, f := range files {
		// Currently only supports flat fs namespace under the plugin directory.
		// TODO: adds support for hierarchical fs namespace.
		if !f.IsDir() && filepath.Base(f.Name())[0] != '.' {
			w.registerPlugin(path.Join(w.path, f.Name()))
		}
	}
	return nil
}

func (w *Watcher) init() error {
	if err := w.createPluginDir(); err != nil {
		return err
	}
	return w.traversePluginDir()
}

func (w *Watcher) registerPlugin(socketPath string) error {
	//TODO: Implement rate limiting to mitigate any DOS kind of attacks.
	glog.V(2).Infof("registerPlugin called for socketPath: %s", socketPath)
	client, conn, err := dial(socketPath)
	if err != nil {
		return fmt.Errorf("dial failed at socket %s, err: %v", socketPath, err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	infoResp, err := client.GetInfo(ctx, &registerapi.InfoRequest{})
	if err != nil {
		glog.Errorf("Failed to get plugin info using RPC GetInfo at socket %s, err: %v", socketPath, err)
		return err
	}
	resp := w.invokeRegistrationCallbackAtHandler(infoResp, socketPath)
	glog.V(2).Infof("call NotifyRegistrationStatus, resp: %v", resp)
	if _, err := client.NotifyRegistrationStatus(ctx, resp); err != nil {
		glog.Errorf("Failed to send registration status at socket %s, err: %v", socketPath, err)
		return err
	}

	if !resp.PluginRegistered {
		glog.Errorf("Registration failed for plugin type: %s, name: %s, socket: %s, err: %s", infoResp.Type, infoResp.Name, socketPath, resp.Error)
		return fmt.Errorf("%s", resp.Error)
	}
	glog.Infof("Successfully registered plugin for plugin type: %s, name: %s, socket: %s", infoResp.Type, infoResp.Name, socketPath)
	return nil
}

func (w *Watcher) invokeRegistrationCallbackAtHandler(infoResp *registerapi.PluginInfo, socketPath string) *registerapi.RegistrationStatus {
	var handlerCbkFn RegisterCallbackFn
	var ok bool
	handlerCbkFn, ok = w.handlers[infoResp.Type]
	if !ok {
		return &registerapi.RegistrationStatus{
			PluginRegistered: false,
			Error:            fmt.Sprintf("No handler found registered for plugin type: %s, socket: %s", infoResp.Type, socketPath),
		}
	}

	var versions []string
	for _, version := range infoResp.SupportedVersions {
		versions = append(versions, version)
	}
	// calls handler callback to verify registration request
	if err := handlerCbkFn(infoResp.Name, versions, socketPath); err != nil {
		return &registerapi.RegistrationStatus{
			PluginRegistered: false,
			Error:            fmt.Sprintf("Plugin registration failed with err: %v", err),
		}
	}

	return &registerapi.RegistrationStatus{PluginRegistered: true}
}

// Start watches for the creation of plugin sockets at the path
func (w *Watcher) Start() error {
	glog.V(2).Infof("Plugin Watcher Start at %s", w.path)
	w.stopCh = make(chan interface{})

	// Creating the directory to be watched if it doesn't exist yet,
	// and walks through the directory to discover the existing plugins.
	if err := w.init(); err != nil {
		return err
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to start plugin watcher, err: %v", err)
	}

	err = watcher.Add(w.path)
	if err != nil {
		watcher.Close()
		return fmt.Errorf("failed to start plugin watcher, err: %v", err)
	}
	w.watcher = watcher
	w.wg.Add(1)
	go func(watcher *fsnotify.Watcher) {
		defer w.wg.Done()
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					err := w.registerPlugin(event.Name)
					if err != nil {
						glog.Errorf("Plugin %s registration failed with error: %v", event.Name, err)
					}
				}
				continue
			case err := <-watcher.Errors:
				//TODO: Handle errors by taking corrective measures
				if err != nil {
					glog.Errorf("Watcher received error: %v", err)
				}
				continue

			case <-w.stopCh:
				watcher.Close()
				break
			}
			break
		}
	}(watcher)
	return nil
}

// Stop stops probing the creation of plugin sockets at the path
func (w *Watcher) Stop() error {
	close(w.stopCh)
	c := make(chan struct{})
	go func() {
		defer close(c)
		w.wg.Wait()
	}()
	select {
	case <-c:
	case <-time.After(10 * time.Second):
		return fmt.Errorf("timeout on stopping watcher")
	}
	return nil
}

// Cleanup cleans the path by removing sockets
func (w *Watcher) Cleanup() error {
	return os.RemoveAll(w.path)
}

// Dial establishes the gRPC communication with the picked up plugin socket. https://godoc.org/google.golang.org/grpc#Dial
func dial(unixSocketPath string) (registerapi.RegistrationClient, *grpc.ClientConn, error) {
	c, err := grpc.Dial(unixSocketPath, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithTimeout(10*time.Second),
		grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}),
	)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial socket %s, err: %v", unixSocketPath, err)
	}

	return registerapi.NewRegistrationClient(c), c, nil
}
