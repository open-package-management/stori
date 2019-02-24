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
)

var repositoryKey contextKey = "repository"

func repoHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		var repo string
		repo, req.URL.Path = shiftPath(req.URL.Path)

		ctx := context.WithValue(req.Context(), repositoryKey, repo)
		req = req.WithContext(ctx)

		if req.URL.Path != "/" {
			resource, _ := shiftPath(req.URL.Path)
			switch resource {
			}
		}

		switch req.Method {
		case "GET":
			handler := getRepoHandler(reg)
			handler.ServeHTTP(w, req)
		case "PUT":
			handler := putRepoHandler(reg)
			handler.ServeHTTP(w, req)
		case "DELETE":
			handler := deleteRepoHandler(reg)
			handler.ServeHTTP(w, req)
		}
	}
	return http.HandlerFunc(fn)
}

func getRepoHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
	return http.HandlerFunc(fn)
}

func putRepoHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
	return http.HandlerFunc(fn)
}

func deleteRepoHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
	return http.HandlerFunc(fn)
}
