// Copyright 2019 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	scheme "github.com/qiankunli/kubeflow-common-v0.1.0/client/clientset/versioned/scheme"
	v1 "github.com/qiankunli/kubeflow-common-v0.1.0/test_job/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TestJobsGetter has a method to return a TestJobInterface.
// A group's client should implement this interface.
type TestJobsGetter interface {
	TestJobs(namespace string) TestJobInterface
}

// TestJobInterface has methods to work with TestJob resources.
type TestJobInterface interface {
	Create(*v1.TestJob) (*v1.TestJob, error)
	Update(*v1.TestJob) (*v1.TestJob, error)
	UpdateStatus(*v1.TestJob) (*v1.TestJob, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.TestJob, error)
	List(opts meta_v1.ListOptions) (*v1.TestJobList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.TestJob, err error)
	TestJobExpansion
}

// testJobs implements TestJobInterface
type testJobs struct {
	client rest.Interface
	ns     string
}

// newTestJobs returns a TestJobs
func newTestJobs(c *KubeflowV1Client, namespace string) *testJobs {
	return &testJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the testJob, and returns the corresponding testJob object, and an error if there is any.
func (c *testJobs) Get(name string, options meta_v1.GetOptions) (result *v1.TestJob, err error) {
	result = &v1.TestJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TestJobs that match those selectors.
func (c *testJobs) List(opts meta_v1.ListOptions) (result *v1.TestJobList, err error) {
	result = &v1.TestJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested testJobs.
func (c *testJobs) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("testjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a testJob and creates it.  Returns the server's representation of the testJob, and an error, if there is any.
func (c *testJobs) Create(testJob *v1.TestJob) (result *v1.TestJob, err error) {
	result = &v1.TestJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("testjobs").
		Body(testJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a testJob and updates it. Returns the server's representation of the testJob, and an error, if there is any.
func (c *testJobs) Update(testJob *v1.TestJob) (result *v1.TestJob, err error) {
	result = &v1.TestJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("testjobs").
		Name(testJob.Name).
		Body(testJob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *testJobs) UpdateStatus(testJob *v1.TestJob) (result *v1.TestJob, err error) {
	result = &v1.TestJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("testjobs").
		Name(testJob.Name).
		SubResource("status").
		Body(testJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the testJob and deletes it. Returns an error if one occurs.
func (c *testJobs) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *testJobs) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched testJob.
func (c *testJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.TestJob, err error) {
	result = &v1.TestJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("testjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
