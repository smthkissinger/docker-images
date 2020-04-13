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

package v1alpha1

import (
	v1alpha1 "github.com/huanwei/rocketmq-operator/pkg/apis/rocketmq/v1alpha1"
	scheme "github.com/huanwei/rocketmq-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BrokerClustersGetter has a method to return a BrokerClusterInterface.
// A group's client should implement this interface.
type BrokerClustersGetter interface {
	BrokerClusters(namespace string) BrokerClusterInterface
}

// BrokerClusterInterface has methods to work with BrokerCluster resources.
type BrokerClusterInterface interface {
	Create(*v1alpha1.BrokerCluster) (*v1alpha1.BrokerCluster, error)
	Update(*v1alpha1.BrokerCluster) (*v1alpha1.BrokerCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BrokerCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.BrokerClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BrokerCluster, err error)
	BrokerClusterExpansion
}

// brokerClusters implements BrokerClusterInterface
type brokerClusters struct {
	client rest.Interface
	ns     string
}

// newBrokerClusters returns a BrokerClusters
func newBrokerClusters(c *ROCKETMQV1alpha1Client, namespace string) *brokerClusters {
	return &brokerClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the brokerCluster, and returns the corresponding brokerCluster object, and an error if there is any.
func (c *brokerClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.BrokerCluster, err error) {
	result = &v1alpha1.BrokerCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("brokerclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BrokerClusters that match those selectors.
func (c *brokerClusters) List(opts v1.ListOptions) (result *v1alpha1.BrokerClusterList, err error) {
	result = &v1alpha1.BrokerClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("brokerclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested brokerClusters.
func (c *brokerClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("brokerclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a brokerCluster and creates it.  Returns the server's representation of the brokerCluster, and an error, if there is any.
func (c *brokerClusters) Create(brokerCluster *v1alpha1.BrokerCluster) (result *v1alpha1.BrokerCluster, err error) {
	result = &v1alpha1.BrokerCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("brokerclusters").
		Body(brokerCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a brokerCluster and updates it. Returns the server's representation of the brokerCluster, and an error, if there is any.
func (c *brokerClusters) Update(brokerCluster *v1alpha1.BrokerCluster) (result *v1alpha1.BrokerCluster, err error) {
	result = &v1alpha1.BrokerCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("brokerclusters").
		Name(brokerCluster.Name).
		Body(brokerCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the brokerCluster and deletes it. Returns an error if one occurs.
func (c *brokerClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("brokerclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *brokerClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("brokerclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched brokerCluster.
func (c *brokerClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BrokerCluster, err error) {
	result = &v1alpha1.BrokerCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("brokerclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}