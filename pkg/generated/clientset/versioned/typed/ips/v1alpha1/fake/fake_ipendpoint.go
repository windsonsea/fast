/*
Copyright The Kubernetes Authors.

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
	"context"

	v1alpha1 "github.com/fast-io/fast/pkg/apis/ips/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIpEndpoints implements IpEndpointInterface
type FakeIpEndpoints struct {
	Fake *FakeSampleV1alpha1
	ns   string
}

var ipendpointsResource = schema.GroupVersionResource{Group: "sample.fast.io", Version: "v1alpha1", Resource: "ipendpoints"}

var ipendpointsKind = schema.GroupVersionKind{Group: "sample.fast.io", Version: "v1alpha1", Kind: "IpEndpoint"}

// Get takes name of the ipEndpoint, and returns the corresponding ipEndpoint object, and an error if there is any.
func (c *FakeIpEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IpEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ipendpointsResource, c.ns, name), &v1alpha1.IpEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpEndpoint), err
}

// List takes label and field selectors, and returns the list of IpEndpoints that match those selectors.
func (c *FakeIpEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IpEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ipendpointsResource, ipendpointsKind, c.ns, opts), &v1alpha1.IpEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IpEndpointList{ListMeta: obj.(*v1alpha1.IpEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.IpEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ipEndpoints.
func (c *FakeIpEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ipendpointsResource, c.ns, opts))

}

// Create takes the representation of a ipEndpoint and creates it.  Returns the server's representation of the ipEndpoint, and an error, if there is any.
func (c *FakeIpEndpoints) Create(ctx context.Context, ipEndpoint *v1alpha1.IpEndpoint, opts v1.CreateOptions) (result *v1alpha1.IpEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ipendpointsResource, c.ns, ipEndpoint), &v1alpha1.IpEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpEndpoint), err
}

// Update takes the representation of a ipEndpoint and updates it. Returns the server's representation of the ipEndpoint, and an error, if there is any.
func (c *FakeIpEndpoints) Update(ctx context.Context, ipEndpoint *v1alpha1.IpEndpoint, opts v1.UpdateOptions) (result *v1alpha1.IpEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ipendpointsResource, c.ns, ipEndpoint), &v1alpha1.IpEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIpEndpoints) UpdateStatus(ctx context.Context, ipEndpoint *v1alpha1.IpEndpoint, opts v1.UpdateOptions) (*v1alpha1.IpEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ipendpointsResource, "status", c.ns, ipEndpoint), &v1alpha1.IpEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpEndpoint), err
}

// Delete takes name of the ipEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeIpEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ipendpointsResource, c.ns, name, opts), &v1alpha1.IpEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIpEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ipendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IpEndpointList{})
	return err
}

// Patch applies the patch and returns the patched ipEndpoint.
func (c *FakeIpEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IpEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ipendpointsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IpEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpEndpoint), err
}
