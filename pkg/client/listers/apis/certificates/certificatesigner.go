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

// Code generated by lister-gen. DO NOT EDIT.

package certificates

import (
	certificates "github.com/fire833/vroute/pkg/apis/certificates"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CertificateSignerLister helps list CertificateSigners.
// All objects returned here must be treated as read-only.
type CertificateSignerLister interface {
	// List lists all CertificateSigners in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*certificates.CertificateSigner, err error)
	// CertificateSigners returns an object that can list and get CertificateSigners.
	CertificateSigners(namespace string) CertificateSignerNamespaceLister
	CertificateSignerListerExpansion
}

// certificateSignerLister implements the CertificateSignerLister interface.
type certificateSignerLister struct {
	indexer cache.Indexer
}

// NewCertificateSignerLister returns a new CertificateSignerLister.
func NewCertificateSignerLister(indexer cache.Indexer) CertificateSignerLister {
	return &certificateSignerLister{indexer: indexer}
}

// List lists all CertificateSigners in the indexer.
func (s *certificateSignerLister) List(selector labels.Selector) (ret []*certificates.CertificateSigner, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*certificates.CertificateSigner))
	})
	return ret, err
}

// CertificateSigners returns an object that can list and get CertificateSigners.
func (s *certificateSignerLister) CertificateSigners(namespace string) CertificateSignerNamespaceLister {
	return certificateSignerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CertificateSignerNamespaceLister helps list and get CertificateSigners.
// All objects returned here must be treated as read-only.
type CertificateSignerNamespaceLister interface {
	// List lists all CertificateSigners in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*certificates.CertificateSigner, err error)
	// Get retrieves the CertificateSigner from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*certificates.CertificateSigner, error)
	CertificateSignerNamespaceListerExpansion
}

// certificateSignerNamespaceLister implements the CertificateSignerNamespaceLister
// interface.
type certificateSignerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CertificateSigners in the indexer for a given namespace.
func (s certificateSignerNamespaceLister) List(selector labels.Selector) (ret []*certificates.CertificateSigner, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*certificates.CertificateSigner))
	})
	return ret, err
}

// Get retrieves the CertificateSigner from the indexer for a given namespace and name.
func (s certificateSignerNamespaceLister) Get(name string) (*certificates.CertificateSigner, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(certificates.Resource("certificatesigner"), name)
	}
	return obj.(*certificates.CertificateSigner), nil
}
