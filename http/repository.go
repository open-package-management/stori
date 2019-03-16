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
	"net/http"

	"github.com/open-package-management/stori/core"
	"github.com/open-package-management/stori/http/internal/pathutil"
)

var repoContextKey contextKey = "repository"

func repositoriesHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {

			var repo string
			repo, req.URL.Path = pathutil.ShiftPath(req.URL.Path)
			ctx := context.WithValue(
				req.Context(),
				repoContextKey,
				repo,
			)

			handler := repositoryHandler(reg)
			handler.ServeHTTP(w, req.WithContext(ctx))
			return
		}

		switch req.Method {
		case "GET":
			handler := getRepositoriesHandler(reg)
			handler.ServeHTTP(w, req)
		case "HEAD":
			handler := headRepositoriesHandler(reg)
			handler.ServeHTTP(w, req)
		default:
			handler := defaultHandler()
			handler.ServeHTTP(w, req)
		}
		return
	}
	return http.HandlerFunc(fn)
}

func repositoryHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {

		if req.URL.Path != "/" {
			var resource string
			resource, req.URL.Path = pathutil.ShiftPath(req.URL.Path)
			switch resource {
			default:
				handler := defaultHandler()
				handler.ServeHTTP(w, req)
			}
			return
		}

		switch req.Method {
		case "GET":
			handler := getRepositoryHandler(reg)
			handler.ServeHTTP(w, req)
		case "HEAD":
			handler := headRepositoryHandler(reg)
			handler.ServeHTTP(w, req)
		case "PUT":
			handler := putRepositoryHandler(reg)
			handler.ServeHTTP(w, req)
		case "DELETE":
			handler := deleteRepositoryHandler(reg)
			handler.ServeHTTP(w, req)
		default:
			handler := defaultHandler()
			handler.ServeHTTP(w, req)
		}
		return
	}
	return http.HandlerFunc(fn)
}

func getRepositoriesHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func headRepositoriesHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func getRepositoryHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func headRepositoryHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func putRepositoryHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func deleteRepositoryHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func repoFromContext(ctx context.Context) string {
	repo := ctx.Value(repoContextKey).(string)
	return repo
}
