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

	"github.com/open-package-management/stori/core"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func metricsHandler(reg core.Registry) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			handler := defaultHandler()
			handler.ServeHTTP(w, req)
		}

		switch req.Method {
		case "GET":
			handler := promhttp.Handler()
			handler.ServeHTTP(w, req)
		default:
			handler := defaultHandler()
			handler.ServeHTTP(w, req)
		}
	}
	return http.HandlerFunc(fn)
}
