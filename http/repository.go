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

var repoContextKey contextKey = "repository"

func baseRepoHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {

			var repo string
			repo, req.URL.Path = shiftPath(req.URL.Path)
			ctx := context.WithValue(
				req.Context(),
				repoContextKey,
				repo,
			)

			handler := repoHandler(reg)
			handler.ServeHTTP(w, req.WithContext(ctx))
			return
		}

		switch req.Method {
		case "GET":
			handler := baseGetRepoHandler(reg)
			handler.ServeHTTP(w, req)
		case "HEAD":
			handler := baseHeadRepoHandler(reg)
			handler.ServeHTTP(w, req)
		default:
			handler := defaultHandler()
			handler.ServeHTTP(w, req)
		}
		return
	}
	return http.HandlerFunc(fn)
}

func repoHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {

		if req.URL.Path != "/" {
			var resource string
			resource, req.URL.Path = shiftPath(req.URL.Path)
			switch resource {
			default:
				handler := defaultHandler()
				handler.ServeHTTP(w, req)
			}
			return
		}

		switch req.Method {
		case "GET":
			handler := getRepoHandler(reg)
			handler.ServeHTTP(w, req)
		case "HEAD":
			handler := headRepoHandler(reg)
			handler.ServeHTTP(w, req)
		case "PUT":
			handler := putRepoHandler(reg)
			handler.ServeHTTP(w, req)
		case "DELETE":
			handler := deleteRepoHandler(reg)
			handler.ServeHTTP(w, req)
		default:
			handler := defaultHandler()
			handler.ServeHTTP(w, req)
		}
		return
	}
	return http.HandlerFunc(fn)
}

func baseGetRepoHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		project := projectFromContext(ctx)
		fmt.Fprintf(w, "LIST Repos Handler\nnamespace: %s\nproject: %s", namespace, project)
	}
	return http.HandlerFunc(fn)
}

func baseHeadRepoHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		project := projectFromContext(ctx)
		fmt.Fprintf(w, "LIST Repos Handler\nnamespace: %s\nproject: %s", namespace, project)
	}
	return http.HandlerFunc(fn)
}

func getRepoHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		project := projectFromContext(ctx)
		repo := repoFromContext(ctx)
		fmt.Fprintf(w, "GET Repo Handler\nnamespace: %s\nproject: %s\nrepo: %s", namespace, project, repo)
	}
	return http.HandlerFunc(fn)
}

func headRepoHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		project := projectFromContext(ctx)
		repo := repoFromContext(ctx)
		fmt.Fprintf(w, "GET Repo Handler\nnamespace: %s\nproject: %s\nrepo: %s", namespace, project, repo)
	}
	return http.HandlerFunc(fn)
}

func putRepoHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		project := projectFromContext(ctx)
		repo := repoFromContext(ctx)
		fmt.Fprintf(w, "PUT Repo Handler\nnamespace: %s\nproject: %s\nrepo: %s", namespace, project, repo)
	}
	return http.HandlerFunc(fn)
}

func deleteRepoHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		ctx := req.Context()
		namespace := namespaceFromContext(ctx)
		project := projectFromContext(ctx)
		repo := repoFromContext(ctx)
		fmt.Fprintf(w, "DELETE Repo Handler\nnamespace: %s\nproject: %s\nrepo: %s", namespace, project, repo)
	}
	return http.HandlerFunc(fn)
}

func repoFromContext(ctx context.Context) string {
	repo := ctx.Value(repoContextKey).(string)
	return repo
}
