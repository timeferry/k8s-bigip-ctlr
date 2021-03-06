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

package v1

import (
	"time"

	v1 "github.com/F5Networks/k8s-bigip-ctlr/config/apis/cis/v1"
	scheme "github.com/F5Networks/k8s-bigip-ctlr/config/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NginxCisConnectorsGetter has a method to return a NginxCisConnectorInterface.
// A group's client should implement this interface.
type NginxCisConnectorsGetter interface {
	NginxCisConnectors(namespace string) NginxCisConnectorInterface
}

// NginxCisConnectorInterface has methods to work with NginxCisConnector resources.
type NginxCisConnectorInterface interface {
	Create(*v1.NginxCisConnector) (*v1.NginxCisConnector, error)
	Update(*v1.NginxCisConnector) (*v1.NginxCisConnector, error)
	UpdateStatus(*v1.NginxCisConnector) (*v1.NginxCisConnector, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.NginxCisConnector, error)
	List(opts metav1.ListOptions) (*v1.NginxCisConnectorList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.NginxCisConnector, err error)
	NginxCisConnectorExpansion
}

// nginxCisConnectors implements NginxCisConnectorInterface
type nginxCisConnectors struct {
	client rest.Interface
	ns     string
}

// newNginxCisConnectors returns a NginxCisConnectors
func newNginxCisConnectors(c *K8sV1Client, namespace string) *nginxCisConnectors {
	return &nginxCisConnectors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nginxCisConnector, and returns the corresponding nginxCisConnector object, and an error if there is any.
func (c *nginxCisConnectors) Get(name string, options metav1.GetOptions) (result *v1.NginxCisConnector, err error) {
	result = &v1.NginxCisConnector{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nginxcisconnectors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NginxCisConnectors that match those selectors.
func (c *nginxCisConnectors) List(opts metav1.ListOptions) (result *v1.NginxCisConnectorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NginxCisConnectorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nginxcisconnectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nginxCisConnectors.
func (c *nginxCisConnectors) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nginxcisconnectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a nginxCisConnector and creates it.  Returns the server's representation of the nginxCisConnector, and an error, if there is any.
func (c *nginxCisConnectors) Create(nginxCisConnector *v1.NginxCisConnector) (result *v1.NginxCisConnector, err error) {
	result = &v1.NginxCisConnector{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nginxcisconnectors").
		Body(nginxCisConnector).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nginxCisConnector and updates it. Returns the server's representation of the nginxCisConnector, and an error, if there is any.
func (c *nginxCisConnectors) Update(nginxCisConnector *v1.NginxCisConnector) (result *v1.NginxCisConnector, err error) {
	result = &v1.NginxCisConnector{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nginxcisconnectors").
		Name(nginxCisConnector.Name).
		Body(nginxCisConnector).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nginxCisConnectors) UpdateStatus(nginxCisConnector *v1.NginxCisConnector) (result *v1.NginxCisConnector, err error) {
	result = &v1.NginxCisConnector{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nginxcisconnectors").
		Name(nginxCisConnector.Name).
		SubResource("status").
		Body(nginxCisConnector).
		Do().
		Into(result)
	return
}

// Delete takes name of the nginxCisConnector and deletes it. Returns an error if one occurs.
func (c *nginxCisConnectors) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nginxcisconnectors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nginxCisConnectors) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nginxcisconnectors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nginxCisConnector.
func (c *nginxCisConnectors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.NginxCisConnector, err error) {
	result = &v1.NginxCisConnector{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nginxcisconnectors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
