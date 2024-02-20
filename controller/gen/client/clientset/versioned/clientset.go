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

package versioned

import (
	"fmt"
	"net/http"

	externalworkloadv1beta1 "github.com/linkerd/linkerd2/controller/gen/client/clientset/versioned/typed/externalworkload/v1beta1"
	linkv1alpha1 "github.com/linkerd/linkerd2/controller/gen/client/clientset/versioned/typed/link/v1alpha1"
	policyv1alpha1 "github.com/linkerd/linkerd2/controller/gen/client/clientset/versioned/typed/policy/v1alpha1"
	policyv1beta3 "github.com/linkerd/linkerd2/controller/gen/client/clientset/versioned/typed/policy/v1beta3"
	serverv1beta2 "github.com/linkerd/linkerd2/controller/gen/client/clientset/versioned/typed/server/v1beta2"
	serverauthorizationv1beta1 "github.com/linkerd/linkerd2/controller/gen/client/clientset/versioned/typed/serverauthorization/v1beta1"
	linkerdv1alpha2 "github.com/linkerd/linkerd2/controller/gen/client/clientset/versioned/typed/serviceprofile/v1alpha2"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ExternalworkloadV1beta1() externalworkloadv1beta1.ExternalworkloadV1beta1Interface
	LinkV1alpha1() linkv1alpha1.LinkV1alpha1Interface
	PolicyV1alpha1() policyv1alpha1.PolicyV1alpha1Interface
	PolicyV1beta3() policyv1beta3.PolicyV1beta3Interface
	ServerV1beta2() serverv1beta2.ServerV1beta2Interface
	ServerauthorizationV1beta1() serverauthorizationv1beta1.ServerauthorizationV1beta1Interface
	LinkerdV1alpha2() linkerdv1alpha2.LinkerdV1alpha2Interface
}

// Clientset contains the clients for groups.
type Clientset struct {
	*discovery.DiscoveryClient
	externalworkloadV1beta1    *externalworkloadv1beta1.ExternalworkloadV1beta1Client
	linkV1alpha1               *linkv1alpha1.LinkV1alpha1Client
	policyV1alpha1             *policyv1alpha1.PolicyV1alpha1Client
	policyV1beta3              *policyv1beta3.PolicyV1beta3Client
	serverV1beta2              *serverv1beta2.ServerV1beta2Client
	serverauthorizationV1beta1 *serverauthorizationv1beta1.ServerauthorizationV1beta1Client
	linkerdV1alpha2            *linkerdv1alpha2.LinkerdV1alpha2Client
}

// ExternalworkloadV1beta1 retrieves the ExternalworkloadV1beta1Client
func (c *Clientset) ExternalworkloadV1beta1() externalworkloadv1beta1.ExternalworkloadV1beta1Interface {
	return c.externalworkloadV1beta1
}

// LinkV1alpha1 retrieves the LinkV1alpha1Client
func (c *Clientset) LinkV1alpha1() linkv1alpha1.LinkV1alpha1Interface {
	return c.linkV1alpha1
}

// PolicyV1alpha1 retrieves the PolicyV1alpha1Client
func (c *Clientset) PolicyV1alpha1() policyv1alpha1.PolicyV1alpha1Interface {
	return c.policyV1alpha1
}

// PolicyV1beta3 retrieves the PolicyV1beta3Client
func (c *Clientset) PolicyV1beta3() policyv1beta3.PolicyV1beta3Interface {
	return c.policyV1beta3
}

// ServerV1beta2 retrieves the ServerV1beta2Client
func (c *Clientset) ServerV1beta2() serverv1beta2.ServerV1beta2Interface {
	return c.serverV1beta2
}

// ServerauthorizationV1beta1 retrieves the ServerauthorizationV1beta1Client
func (c *Clientset) ServerauthorizationV1beta1() serverauthorizationv1beta1.ServerauthorizationV1beta1Interface {
	return c.serverauthorizationV1beta1
}

// LinkerdV1alpha2 retrieves the LinkerdV1alpha2Client
func (c *Clientset) LinkerdV1alpha2() linkerdv1alpha2.LinkerdV1alpha2Interface {
	return c.linkerdV1alpha2
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.externalworkloadV1beta1, err = externalworkloadv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.linkV1alpha1, err = linkv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.policyV1alpha1, err = policyv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.policyV1beta3, err = policyv1beta3.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.serverV1beta2, err = serverv1beta2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.serverauthorizationV1beta1, err = serverauthorizationv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.linkerdV1alpha2, err = linkerdv1alpha2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.externalworkloadV1beta1 = externalworkloadv1beta1.New(c)
	cs.linkV1alpha1 = linkv1alpha1.New(c)
	cs.policyV1alpha1 = policyv1alpha1.New(c)
	cs.policyV1beta3 = policyv1beta3.New(c)
	cs.serverV1beta2 = serverv1beta2.New(c)
	cs.serverauthorizationV1beta1 = serverauthorizationv1beta1.New(c)
	cs.linkerdV1alpha2 = linkerdv1alpha2.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
