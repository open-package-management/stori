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

var projectContextKey contextKey = "project"

func baseProjectHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {

			var project string
			project, req.URL.Path = shiftPath(req.URL.Path)
			ctx := context.WithValue(
				req.Context(),
				projectContextKey,
				project,
			)

			handler := projectHandler(reg)
			handler.ServeHTTP(w, req.WithContext(ctx))
			return
		}

		switch req.Method {
		case "GET":
			handler := listProjectsHandler(reg)
			handler.ServeHTTP(w, req)
		}
		return
	}
	return http.HandlerFunc(fn)
}

func projectHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			var resource string
			resource, req.URL.Path = shiftPath(req.URL.Path)
			switch resource {
			case "repositories":
				handler := baseRepoHandler(reg)
				handler.ServeHTTP(w, req)
			default:
				handler := notFoundHandler()
				handler.ServeHTTP(w, req)
			}
			return
		}

		switch req.Method {
		case "GET":
			handler := getProjectHandler(reg)
			handler.ServeHTTP(w, req)
		case "PUT":
			handler := putProjectHandler(reg)
			handler.ServeHTTP(w, req)
		case "DELETE":
			handler := deleteProjectHandler(reg)
			handler.ServeHTTP(w, req)
		}
		return
	}
	return http.HandlerFunc(fn)
}

func listProjectsHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		fmt.Fprintf(w, "LIST Projects Handler\nnamespace: %s", namespace)
	}
	return http.HandlerFunc(fn)
}

func getProjectHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		project := projectFromContext(ctx)
		fmt.Fprintf(w, "GET Project Handler\nnamespace: %s\nproject: %s", namespace, project)
	}
	return http.HandlerFunc(fn)
}

func putProjectHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		project := projectFromContext(ctx)
		fmt.Fprintf(w, "PUT Project Handler\nnamespace: %s\nproject: %s", namespace, project)
	}
	return http.HandlerFunc(fn)
}

func deleteProjectHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		project := projectFromContext(ctx)
		fmt.Fprintf(w, "DELETE Project Handler\nnamespace: %s\nproject: %s", namespace, project)
	}
	return http.HandlerFunc(fn)
}

func projectFromContext(ctx context.Context) string {
	project, _ := ctx.Value(projectContextKey).(string)
	return project
}
