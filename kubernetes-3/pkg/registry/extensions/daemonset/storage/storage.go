/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storage

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-3/pkg/api"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-3/pkg/apis/extensions"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-3/pkg/registry/cachesize"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-3/pkg/registry/extensions/daemonset"
)

// rest implements a RESTStorage for DaemonSets
type REST struct {
	*genericregistry.Store
}

// NewREST returns a RESTStorage object that will work against DaemonSets.
func NewREST(optsGetter generic.RESTOptionsGetter) (*REST, *StatusREST) {
	store := &genericregistry.Store{
		Copier:      api.Scheme,
		NewFunc:     func() runtime.Object { return &extensions.DaemonSet{} },
		NewListFunc: func() runtime.Object { return &extensions.DaemonSetList{} },
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*extensions.DaemonSet).Name, nil
		},
		PredicateFunc:     daemonset.MatchDaemonSet,
		QualifiedResource: extensions.Resource("daemonsets"),
		WatchCacheSize:    cachesize.GetWatchCacheSizeByResource("daemonsets"),

		CreateStrategy: daemonset.Strategy,
		UpdateStrategy: daemonset.Strategy,
		DeleteStrategy: daemonset.Strategy,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: daemonset.GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err) // TODO: Propagate error up
	}

	statusStore := *store
	statusStore.UpdateStrategy = daemonset.StatusStrategy

	return &REST{store}, &StatusREST{store: &statusStore}
}

// Implement ShortNamesProvider
var _ rest.ShortNamesProvider = &REST{}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"ds"}
}

// StatusREST implements the REST endpoint for changing the status of a daemonset
type StatusREST struct {
	store *genericregistry.Store
}

func (r *StatusREST) New() runtime.Object {
	return &extensions.DaemonSet{}
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *StatusREST) Get(ctx genericapirequest.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return r.store.Get(ctx, name, options)
}

// Update alters the status subset of an object.
func (r *StatusREST) Update(ctx genericapirequest.Context, name string, objInfo rest.UpdatedObjectInfo) (runtime.Object, bool, error) {
	return r.store.Update(ctx, name, objInfo)
}
