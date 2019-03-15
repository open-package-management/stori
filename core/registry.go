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

package core

import (
	"time"
)

type Registry struct{}

type Namespace struct {
	name        string
	status      string
	quotas      NamespaceQuotas
	created     *time.Time
	deleted     *time.Time
	labels      map[string]string
	annotations map[string]string
}

type NamespaceQuotas struct {
	Storage      Quota
	Repositories Quota
	Projects     Quota
}

type Quota struct {
	limit uint64
	used  uint64
}

func (q *Quota) Limit() uint64 {
	return q.limit
}

func (q *Quota) Used() uint64 {
	return q.limit
}

func (ns *Namespace) Name() string {
	return ns.name
}

func (ns *Namespace) Status() string {
	return ns.status
}

func (ns *Namespace) Quotas() NamespaceQuotas {
	return ns.quotas
}

func (ns *Namespace) Labels() map[string]string {
	return ns.labels
}

func (ns *Namespace) Annotations() map[string]string {
	return ns.labels
}

func (ns *Namespace) CreatedTimestamp() *time.Time {
	return ns.created
}

func (ns *Namespace) DeletedTimestamp() *time.Time {
	return ns.deleted
}

// NamespaceList retrieves all namespaces in a registry.
func (reg *Registry) NamespaceList() (*[]Namespace, error) {
	return nil, nil
}
