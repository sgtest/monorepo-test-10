// +build windows

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

package cm

import (
	"github.com/golang/glog"

	"k8s.io/client-go/tools/record"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-12/pkg/api/v1"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-12/pkg/kubelet/cadvisor"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-12/pkg/util/mount"
)

type containerManagerImpl struct {
	containerManagerStub
}

var _ ContainerManager = &containerManagerImpl{}

func (cm *containerManagerImpl) Start(_ *v1.Node, _ ActivePodsFunc) error {
	glog.V(2).Infof("Starting Windows stub container manager")
	return nil
}

func NewContainerManager(mountUtil mount.Interface, cadvisorInterface cadvisor.Interface, nodeConfig NodeConfig, failSwapOn bool, recorder record.EventRecorder) (ContainerManager, error) {
	return &containerManagerImpl{}, nil
}
