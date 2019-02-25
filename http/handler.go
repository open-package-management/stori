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

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

type contextKey string

// Handler is the handler used to serve the registry.
func Handler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		var resource string
		resource, req.URL.Path = shiftPath(req.URL.Path)

		switch resource {
		case "namespaces":
			handler := baseNamespaceHandler(reg)
			handler.ServeHTTP(w, req)
		default:
			handler := defaultHandler()
			handler.ServeHTTP(w, req)
		}
		return
	}
	return http.HandlerFunc(fn)
}

func defaultHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET", "HEAD":
			handler := notFoundHandler()
			handler.ServeHTTP(w, req)
		default:
			handler := notImplementedHandler()
			handler.ServeHTTP(w, req)
		}
	}
	return http.HandlerFunc(fn)
}

func notFoundHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}
	return http.HandlerFunc(fn)
}

func notImplementedHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
	return http.HandlerFunc(fn)
}
