/*
Copyright 2019 Red Hat, Inc.

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
	autoscalingv1 "github.com/openshift/cluster-resource-override-admission-operator/pkg/apis/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterResourceOverrides implements ClusterResourceOverrideInterface
type FakeClusterResourceOverrides struct {
	Fake *FakeAutoscalingV1
}

var clusterresourceoverridesResource = schema.GroupVersionResource{Group: "autoscaling.openshift.io", Version: "v1", Resource: "clusterresourceoverrides"}

var clusterresourceoverridesKind = schema.GroupVersionKind{Group: "autoscaling.openshift.io", Version: "v1", Kind: "ClusterResourceOverride"}

// Get takes name of the clusterResourceOverride, and returns the corresponding clusterResourceOverride object, and an error if there is any.
func (c *FakeClusterResourceOverrides) Get(name string, options v1.GetOptions) (result *autoscalingv1.ClusterResourceOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterresourceoverridesResource, name), &autoscalingv1.ClusterResourceOverride{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.ClusterResourceOverride), err
}

// List takes label and field selectors, and returns the list of ClusterResourceOverrides that match those selectors.
func (c *FakeClusterResourceOverrides) List(opts v1.ListOptions) (result *autoscalingv1.ClusterResourceOverrideList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterresourceoverridesResource, clusterresourceoverridesKind, opts), &autoscalingv1.ClusterResourceOverrideList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &autoscalingv1.ClusterResourceOverrideList{ListMeta: obj.(*autoscalingv1.ClusterResourceOverrideList).ListMeta}
	for _, item := range obj.(*autoscalingv1.ClusterResourceOverrideList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterResourceOverrides.
func (c *FakeClusterResourceOverrides) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterresourceoverridesResource, opts))
}

// Create takes the representation of a clusterResourceOverride and creates it.  Returns the server's representation of the clusterResourceOverride, and an error, if there is any.
func (c *FakeClusterResourceOverrides) Create(clusterResourceOverride *autoscalingv1.ClusterResourceOverride) (result *autoscalingv1.ClusterResourceOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterresourceoverridesResource, clusterResourceOverride), &autoscalingv1.ClusterResourceOverride{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.ClusterResourceOverride), err
}

// Update takes the representation of a clusterResourceOverride and updates it. Returns the server's representation of the clusterResourceOverride, and an error, if there is any.
func (c *FakeClusterResourceOverrides) Update(clusterResourceOverride *autoscalingv1.ClusterResourceOverride) (result *autoscalingv1.ClusterResourceOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterresourceoverridesResource, clusterResourceOverride), &autoscalingv1.ClusterResourceOverride{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.ClusterResourceOverride), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterResourceOverrides) UpdateStatus(clusterResourceOverride *autoscalingv1.ClusterResourceOverride) (*autoscalingv1.ClusterResourceOverride, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterresourceoverridesResource, "status", clusterResourceOverride), &autoscalingv1.ClusterResourceOverride{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.ClusterResourceOverride), err
}

// Delete takes name of the clusterResourceOverride and deletes it. Returns an error if one occurs.
func (c *FakeClusterResourceOverrides) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterresourceoverridesResource, name), &autoscalingv1.ClusterResourceOverride{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterResourceOverrides) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterresourceoverridesResource, listOptions)

	_, err := c.Fake.Invokes(action, &autoscalingv1.ClusterResourceOverrideList{})
	return err
}

// Patch applies the patch and returns the patched clusterResourceOverride.
func (c *FakeClusterResourceOverrides) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *autoscalingv1.ClusterResourceOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterresourceoverridesResource, name, pt, data, subresources...), &autoscalingv1.ClusterResourceOverride{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.ClusterResourceOverride), err
}