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

package fake

import (
	clientset "pkg/client/clientset/versioned"
	dnsv1alpha1 "pkg/client/clientset/versioned/typed/dns/v1alpha1"
	fakednsv1alpha1 "pkg/client/clientset/versioned/typed/dns/v1alpha1/fake"
	firewallv1alpha1 "pkg/client/clientset/versioned/typed/firewall/v1alpha1"
	fakefirewallv1alpha1 "pkg/client/clientset/versioned/typed/firewall/v1alpha1/fake"
	netv1alpha1 "pkg/client/clientset/versioned/typed/net/v1alpha1"
	fakenetv1alpha1 "pkg/client/clientset/versioned/typed/net/v1alpha1/fake"
	servicesv1alpha1 "pkg/client/clientset/versioned/typed/services/v1alpha1"
	fakeservicesv1alpha1 "pkg/client/clientset/versioned/typed/services/v1alpha1/fake"
	vpnv1alpha1 "pkg/client/clientset/versioned/typed/vpn/v1alpha1"
	fakevpnv1alpha1 "pkg/client/clientset/versioned/typed/vpn/v1alpha1/fake"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// DnsV1alpha1 retrieves the DnsV1alpha1Client
func (c *Clientset) DnsV1alpha1() dnsv1alpha1.DnsV1alpha1Interface {
	return &fakednsv1alpha1.FakeDnsV1alpha1{Fake: &c.Fake}
}

// FirewallV1alpha1 retrieves the FirewallV1alpha1Client
func (c *Clientset) FirewallV1alpha1() firewallv1alpha1.FirewallV1alpha1Interface {
	return &fakefirewallv1alpha1.FakeFirewallV1alpha1{Fake: &c.Fake}
}

// NetV1alpha1 retrieves the NetV1alpha1Client
func (c *Clientset) NetV1alpha1() netv1alpha1.NetV1alpha1Interface {
	return &fakenetv1alpha1.FakeNetV1alpha1{Fake: &c.Fake}
}

// ServicesV1alpha1 retrieves the ServicesV1alpha1Client
func (c *Clientset) ServicesV1alpha1() servicesv1alpha1.ServicesV1alpha1Interface {
	return &fakeservicesv1alpha1.FakeServicesV1alpha1{Fake: &c.Fake}
}

// VpnV1alpha1 retrieves the VpnV1alpha1Client
func (c *Clientset) VpnV1alpha1() vpnv1alpha1.VpnV1alpha1Interface {
	return &fakevpnv1alpha1.FakeVpnV1alpha1{Fake: &c.Fake}
}
