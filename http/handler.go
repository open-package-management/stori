// Copyright © 2019 Stori Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"net/http"
	"path"
	"strings"

	"github.com/open-package-management/stori/core"
)

type contextKey string

// StoriHandler is the handler used to serve the registry.
func StoriHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		var resource string
		resource, req.URL.Path = shiftPath(req.URL.Path)

		switch resource {
		case "namespaces":
			handler := namespaceHandler(reg)
			handler.ServeHTTP(w, req)
		}
	}
	return http.HandlerFunc(fn)
}

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}