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
	"context"
	"fmt"
	"net/http"

	"github.com/open-package-management/stori/core"
)

var namespaceKey contextKey = "namespace"

func namespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		var namespace string
		namespace, req.URL.Path = shiftPath(req.URL.Path)

		ctx := context.WithValue(req.Context(), namespaceKey, namespace)
		req = req.WithContext(ctx)
		if req.URL.Path != "/" {
			resource, _ := shiftPath(req.URL.Path)
			switch resource {
			case "projects":
				fmt.Println("project handler activated")
				handler := projectHandler(reg)
				handler.ServeHTTP(w, req)
				return
			default:
				fmt.Println("not found handler activated")
				handler := notFoundHandler()
				handler.ServeHTTP(w, req)
				return
			}
		}

		switch req.Method {
		case "GET":
			fmt.Println("get namespace handler activated")
			handler := getNamespaceHandler(reg)
			handler.ServeHTTP(w, req)
			return
		case "PUT":
			fmt.Println("put namespace handler activated")
			handler := putNamespaceHandler(reg)
			handler.ServeHTTP(w, req)
			return
		case "DELETE":
			fmt.Println("delete namespace handler activated")
			handler := deleteNamespaceHandler(reg)
			handler.ServeHTTP(w, req)
			return
		}

	}
	return http.HandlerFunc(fn)

}

func getNamespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
	return http.HandlerFunc(fn)
}

func putNamespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
	return http.HandlerFunc(fn)
}

func deleteNamespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
	return http.HandlerFunc(fn)
}
