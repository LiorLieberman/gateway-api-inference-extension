/*
Copyright 2025 The Kubernetes Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha2 "sigs.k8s.io/gateway-api-inference-extension/client-go/clientset/versioned/typed/api/v1alpha2"
)

type FakeInferenceV1alpha2 struct {
	*testing.Fake
}

func (c *FakeInferenceV1alpha2) InferenceModels(namespace string) v1alpha2.InferenceModelInterface {
	return newFakeInferenceModels(c, namespace)
}

func (c *FakeInferenceV1alpha2) InferencePools(namespace string) v1alpha2.InferencePoolInterface {
	return newFakeInferencePools(c, namespace)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeInferenceV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
