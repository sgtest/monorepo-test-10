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

package daemonset

import (
	"testing"

	"k8s.io/apiserver/pkg/registry/rest"
	_ "github.com/sourcegraph/monorepo-test-1/kubernetes-1/pkg/api"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-1/pkg/api/testapi"
	apitesting "github.com/sourcegraph/monorepo-test-1/kubernetes-1/pkg/api/testing"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-1/pkg/apis/extensions"
)

func TestSelectableFieldLabelConversions(t *testing.T) {
	apitesting.TestSelectableFieldLabelConversionsOfKind(t,
		testapi.Extensions.GroupVersion().String(),
		"DaemonSet",
		DaemonSetToSelectableFields(&extensions.DaemonSet{}),
		nil,
	)
}

func TestDefaultGarbageCollectionPolicy(t *testing.T) {
	// Make sure we correctly implement the interface.
	// Otherwise a typo could silently change the default.
	var gcds rest.GarbageCollectionDeleteStrategy = Strategy
	if got, want := gcds.DefaultGarbageCollectionPolicy(), rest.OrphanDependents; got != want {
		t.Errorf("DefaultGarbageCollectionPolicy() = %#v, want %#v", got, want)
	}
}
