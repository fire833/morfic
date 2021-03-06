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

package internalversion

import (
	"net/http"

	"github.com/fire833/morfic/pkg/client/clientset/internalversion/scheme"
	rest "k8s.io/client-go/rest"
)

type CertificatesInterface interface {
	RESTClient() rest.Interface
	CertificatesGetter
	CertificateListsGetter
	CertificateSignersGetter
	CertificateSignerListsGetter
	CertificateSigningRequestsGetter
	CertificateSigningRequestListsGetter
}

// CertificatesClient is used to interact with features provided by the certificates.morfic.io group.
type CertificatesClient struct {
	restClient rest.Interface
}

func (c *CertificatesClient) Certificates(namespace string) CertificateInterface {
	return newCertificates(c, namespace)
}

func (c *CertificatesClient) CertificateLists(namespace string) CertificateListInterface {
	return newCertificateLists(c, namespace)
}

func (c *CertificatesClient) CertificateSigners(namespace string) CertificateSignerInterface {
	return newCertificateSigners(c, namespace)
}

func (c *CertificatesClient) CertificateSignerLists(namespace string) CertificateSignerListInterface {
	return newCertificateSignerLists(c, namespace)
}

func (c *CertificatesClient) CertificateSigningRequests(namespace string) CertificateSigningRequestInterface {
	return newCertificateSigningRequests(c, namespace)
}

func (c *CertificatesClient) CertificateSigningRequestLists(namespace string) CertificateSigningRequestListInterface {
	return newCertificateSigningRequestLists(c, namespace)
}

// NewForConfig creates a new CertificatesClient for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*CertificatesClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new CertificatesClient for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*CertificatesClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &CertificatesClient{client}, nil
}

// NewForConfigOrDie creates a new CertificatesClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CertificatesClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CertificatesClient for the given RESTClient.
func New(c rest.Interface) *CertificatesClient {
	return &CertificatesClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("certificates.morfic.io")[0].Group {
		gv := scheme.Scheme.PrioritizedVersionsForGroup("certificates.morfic.io")[0]
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs

	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CertificatesClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
