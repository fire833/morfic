/*
*	Copyright (C) 2022  Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	certificatesv1alpha1 "github.com/fire833/morfic/pkg/client/clientset/versioned/typed/certificates/v1alpha1"
	dnsv1alpha1 "github.com/fire833/morfic/pkg/client/clientset/versioned/typed/dns/v1alpha1"
	firewallv1alpha1 "github.com/fire833/morfic/pkg/client/clientset/versioned/typed/firewall/v1alpha1"
	netv1alpha1 "github.com/fire833/morfic/pkg/client/clientset/versioned/typed/net/v1alpha1"
	servicesv1alpha1 "github.com/fire833/morfic/pkg/client/clientset/versioned/typed/services/v1alpha1"
	vpnv1alpha1 "github.com/fire833/morfic/pkg/client/clientset/versioned/typed/vpn/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	CertificatesV1alpha1() certificatesv1alpha1.CertificatesV1alpha1Interface
	DnsV1alpha1() dnsv1alpha1.DnsV1alpha1Interface
	FirewallV1alpha1() firewallv1alpha1.FirewallV1alpha1Interface
	NetV1alpha1() netv1alpha1.NetV1alpha1Interface
	ServicesV1alpha1() servicesv1alpha1.ServicesV1alpha1Interface
	VpnV1alpha1() vpnv1alpha1.VpnV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	certificatesV1alpha1 *certificatesv1alpha1.CertificatesV1alpha1Client
	dnsV1alpha1          *dnsv1alpha1.DnsV1alpha1Client
	firewallV1alpha1     *firewallv1alpha1.FirewallV1alpha1Client
	netV1alpha1          *netv1alpha1.NetV1alpha1Client
	servicesV1alpha1     *servicesv1alpha1.ServicesV1alpha1Client
	vpnV1alpha1          *vpnv1alpha1.VpnV1alpha1Client
}

// CertificatesV1alpha1 retrieves the CertificatesV1alpha1Client
func (c *Clientset) CertificatesV1alpha1() certificatesv1alpha1.CertificatesV1alpha1Interface {
	return c.certificatesV1alpha1
}

// DnsV1alpha1 retrieves the DnsV1alpha1Client
func (c *Clientset) DnsV1alpha1() dnsv1alpha1.DnsV1alpha1Interface {
	return c.dnsV1alpha1
}

// FirewallV1alpha1 retrieves the FirewallV1alpha1Client
func (c *Clientset) FirewallV1alpha1() firewallv1alpha1.FirewallV1alpha1Interface {
	return c.firewallV1alpha1
}

// NetV1alpha1 retrieves the NetV1alpha1Client
func (c *Clientset) NetV1alpha1() netv1alpha1.NetV1alpha1Interface {
	return c.netV1alpha1
}

// ServicesV1alpha1 retrieves the ServicesV1alpha1Client
func (c *Clientset) ServicesV1alpha1() servicesv1alpha1.ServicesV1alpha1Interface {
	return c.servicesV1alpha1
}

// VpnV1alpha1 retrieves the VpnV1alpha1Client
func (c *Clientset) VpnV1alpha1() vpnv1alpha1.VpnV1alpha1Interface {
	return c.vpnV1alpha1
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
	cs.certificatesV1alpha1, err = certificatesv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.dnsV1alpha1, err = dnsv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.firewallV1alpha1, err = firewallv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.netV1alpha1, err = netv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.servicesV1alpha1, err = servicesv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.vpnV1alpha1, err = vpnv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
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
	cs.certificatesV1alpha1 = certificatesv1alpha1.New(c)
	cs.dnsV1alpha1 = dnsv1alpha1.New(c)
	cs.firewallV1alpha1 = firewallv1alpha1.New(c)
	cs.netV1alpha1 = netv1alpha1.New(c)
	cs.servicesV1alpha1 = servicesv1alpha1.New(c)
	cs.vpnV1alpha1 = vpnv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
