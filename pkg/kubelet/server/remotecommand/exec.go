/*
Copyright 2016 The Kubernetes Authors.

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

package remotecommand

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/types"
	utilexec "k8s.io/kubernetes/pkg/util/exec"
	"k8s.io/kubernetes/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/util/term"
)

// Executor knows how to execute a command in a container in a pod.
type Executor interface {
	// ExecInContainer executes a command in a container in the pod, copying data
	// between in/out/err and the container's stdin/stdout/stderr.
	ExecInContainer(name string, uid types.UID, container string, cmd []string, in io.Reader, out, err io.WriteCloser, tty bool, resize <-chan term.Size) error
}

// ServeExec handles requests to execute a command in a container. After
// creating/receiving the required streams, it delegates the actual execution
// to the executor.
func ServeExec(w http.ResponseWriter, req *http.Request, executor Executor, podName string, uid types.UID, container string, idleTimeout, streamCreationTimeout time.Duration, supportedProtocols []string) {
	ctx, ok := createStreams(req, w, supportedProtocols, idleTimeout, streamCreationTimeout)
	if !ok {
		// error is handled by createStreams
		return
	}
	defer ctx.conn.Close()

	cmd := req.URL.Query()[api.ExecCommandParamm]

	err := executor.ExecInContainer(podName, uid, container, cmd, ctx.stdinStream, ctx.stdoutStream, ctx.stderrStream, ctx.tty, ctx.resizeChan)
	if err != nil {
		msg := fmt.Sprintf("error executing command in container: %v", err)
		var rc *int
		if exitErr, ok := err.(utilexec.ExitError); ok {
			msg = exitErr.Error()
			if exitErr.Exited() {
				c := exitErr.ExitStatus()
				rc = &c
			}
		}
		runtime.HandleError(errors.New(msg))
		ctx.writeError(msg, rc)
	}
}
