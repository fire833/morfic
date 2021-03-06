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
	certificates "github.com/fire833/morfic/pkg/apis/certificates"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CertificateSigningRequestListLister helps list CertificateSigningRequestLists.
// All objects returned here must be treated as read-only.
type CertificateSigningRequestListLister interface {
	// List lists all CertificateSigningRequestLists in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*certificates.CertificateSigningRequestList, err error)
	// CertificateSigningRequestLists returns an object that can list and get CertificateSigningRequestLists.
	CertificateSigningRequestLists(namespace string) CertificateSigningRequestListNamespaceLister
	CertificateSigningRequestListListerExpansion
}

// certificateSigningRequestListLister implements the CertificateSigningRequestListLister interface.
type certificateSigningRequestListLister struct {
	indexer cache.Indexer
}

// NewCertificateSigningRequestListLister returns a new CertificateSigningRequestListLister.
func NewCertificateSigningRequestListLister(indexer cache.Indexer) CertificateSigningRequestListLister {
	return &certificateSigningRequestListLister{indexer: indexer}
}

// List lists all CertificateSigningRequestLists in the indexer.
func (s *certificateSigningRequestListLister) List(selector labels.Selector) (ret []*certificates.CertificateSigningRequestList, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*certificates.CertificateSigningRequestList))
	})
	return ret, err
}

// CertificateSigningRequestLists returns an object that can list and get CertificateSigningRequestLists.
func (s *certificateSigningRequestListLister) CertificateSigningRequestLists(namespace string) CertificateSigningRequestListNamespaceLister {
	return certificateSigningRequestListNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CertificateSigningRequestListNamespaceLister helps list and get CertificateSigningRequestLists.
// All objects returned here must be treated as read-only.
type CertificateSigningRequestListNamespaceLister interface {
	// List lists all CertificateSigningRequestLists in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*certificates.CertificateSigningRequestList, err error)
	// Get retrieves the CertificateSigningRequestList from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*certificates.CertificateSigningRequestList, error)
	CertificateSigningRequestListNamespaceListerExpansion
}

// certificateSigningRequestListNamespaceLister implements the CertificateSigningRequestListNamespaceLister
// interface.
type certificateSigningRequestListNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CertificateSigningRequestLists in the indexer for a given namespace.
func (s certificateSigningRequestListNamespaceLister) List(selector labels.Selector) (ret []*certificates.CertificateSigningRequestList, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*certificates.CertificateSigningRequestList))
	})
	return ret, err
}

// Get retrieves the CertificateSigningRequestList from the indexer for a given namespace and name.
func (s certificateSigningRequestListNamespaceLister) Get(name string) (*certificates.CertificateSigningRequestList, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(certificates.Resource("certificatesigningrequestlist"), name)
	}
	return obj.(*certificates.CertificateSigningRequestList), nil
}
