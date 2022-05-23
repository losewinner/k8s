//go:build windows
// +build windows

/*
Copyright 2022 The Kubernetes Authors.

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

package filesystem

import (
	"path/filepath"
)

// IsAbs returns wether the given path is absolute or not.
// On Windows, filepath.IsAbs will not return True for paths prefixed with a slash, even
// though they can be used as absolute paths (https://docs.microsoft.com/en-us/dotnet/standard/io/file-path-formats).
func IsAbs(path string) bool {
	return filepath.IsAbs(path) || (len(path) > 0 && (path[0] == '\\' || path[0] == '/'))
}
