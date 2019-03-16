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
	"encoding/json"
	"net/http"

	"github.com/open-package-management/stori/core"
	"github.com/open-package-management/stori/http/internal/pathutil"

	specV1 "github.com/open-package-management/go-specs/v1"
)

var namespaceContextKey contextKey = "namespace"

func namespacesHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			var namespace string
			namespace, req.URL.Path = pathutil.ShiftPath(req.URL.Path)
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
			handler := getNamespacesHandler(reg)
			handler.ServeHTTP(w, req)
		case "HEAD":
			handler := headNamespacesHandler(reg)
			handler.ServeHTTP(w, req)
		default:
			handler := defaultHandler()
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
			resource, req.URL.Path = pathutil.ShiftPath(req.URL.Path)
			switch resource {
			case "projects":
				handler := projectsHandler(reg)
				handler.ServeHTTP(w, req)
			default:
				handler := defaultHandler()
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
			handler := defaultHandler()
			handler.ServeHTTP(w, req)
		}
		return
	}
	return http.HandlerFunc(fn)
}

func getNamespacesHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		namespaces, err := reg.NamespaceList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		list, err := formatNamespaceList(namespaces, specV1.MediaTypeNamespace)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		enc := json.NewEncoder(w)
		if err := enc.Encode(list); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
	return http.HandlerFunc(fn)
}

func formatNamespaceList(namespaces *[]core.Namespace, mediaType string) (interface{}, error) {
	switch mediaType {
	case specV1.MediaTypeNamespace:
		list := formatNamespaceListV1(namespaces)
		return list, nil
	default:
		return nil, nil
	}
}

func formatNamespaceListV1(namespaces *[]core.Namespace) specV1.NamespaceList {
	var list specV1.NamespaceList
	for idx, ns := range *namespaces {
		quotas := ns.Quotas()
		projects := quotas.Projects
		storage := quotas.Storage
		repos := quotas.Repositories

		nsQuotas := specV1.NamespaceQuotas{
			Storage: specV1.QuotaDescriptor{
				Limit: storage.Limit(),
				Used:  storage.Used(),
			},
			Repositories: specV1.QuotaDescriptor{
				Limit: repos.Limit(),
				Used:  repos.Used(),
			},
			Projects: specV1.QuotaDescriptor{
				Limit: projects.Limit(),
				Used:  projects.Used(),
			},
		}

		list.Namespaces[idx] = specV1.Namespace{
			Name:        ns.Name(),
			Status:      ns.Status(),
			Quotas:      nsQuotas,
			Labels:      ns.Labels(),
			Annotations: ns.Annotations(),
			Created:     *ns.CreatedTimestamp(),
			Deleted:     *ns.DeletedTimestamp(),
		}
	}
	return list
}

func headNamespacesHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func getNamespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
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
	}
	return http.HandlerFunc(fn)
}

func deleteNamespaceHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

func namespaceFromContext(ctx context.Context) string {
	namespace := ctx.Value(namespaceContextKey).(string)
	return namespace
}
