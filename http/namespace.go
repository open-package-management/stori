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

var namespaceContextKey contextKey = "namespace"

func baseNamespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {

			var namespace string
			namespace, req.URL.Path = shiftPath(req.URL.Path)
			ctx := context.WithValue(
				req.Context(),
				namespaceContextKey,
				namespace,
			)

			handler := namespaceHandler(reg)
			handler.ServeHTTP(w, req.WithContext(ctx))
			return
		}

		switch req.Method {
		case "GET":
			handler := baseGetNamespaceHandler(reg)
			handler.ServeHTTP(w, req)
		case "HEAD":
			handler := baseHeadNamespaceHandler(reg)
			handler.ServeHTTP(w, req)
		default:
			handler := notImplementedHandler()
			handler.ServeHTTP(w, req)
		}
		return

	}
	return http.HandlerFunc(fn)

}

func namespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {

		if req.URL.Path != "/" {
			var resource string
			resource, req.URL.Path = shiftPath(req.URL.Path)
			switch resource {
			case "projects":
				handler := projectHandler(reg)
				handler.ServeHTTP(w, req)
			default:
				handler := notFoundHandler()
				handler.ServeHTTP(w, req)
			}
			return
		}

		switch req.Method {
		case "GET":
			handler := getNamespaceHandler(reg)
			handler.ServeHTTP(w, req)
		case "HEAD":
			handler := headNamespaceHandler(reg)
			handler.ServeHTTP(w, req)
		case "PUT":
			handler := putNamespaceHandler(reg)
			handler.ServeHTTP(w, req)
		case "DELETE":
			handler := deleteNamespaceHandler(reg)
			handler.ServeHTTP(w, req)
		default:
			handler := notFoundHandler()
			handler.ServeHTTP(w, req)
		}
		return
	}
	return http.HandlerFunc(fn)
}

func baseGetNamespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "LIST Namespaces Handler")
	}
	return http.HandlerFunc(fn)
}

func baseHeadNamespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func getNamespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		fmt.Fprintf(w, "GET Namespace Handler\n namespace: %s", namespace)
	}
	return http.HandlerFunc(fn)
}

func headNamespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func putNamespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		fmt.Fprintf(w, "PUT Namespace Handler\n namespace: %s", namespace)
	}
	return http.HandlerFunc(fn)
}

func deleteNamespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		fmt.Fprintf(w, "DELETE Namespace Handler\n namespace: %s", namespace)
	}
	return http.HandlerFunc(fn)
}

func namespaceFromContext(ctx context.Context) string {
	namespace := ctx.Value(namespaceContextKey).(string)
	return namespace
}
