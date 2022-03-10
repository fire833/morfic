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

// Code generated by informer-gen. DO NOT EDIT.

package certificates

import (
	"context"
	internalinterfaces "informers/internalinterfaces"
	time "time"

	apiscertificates "github.com/fire833/vroute/pkg/apis/certificates"
	clientset "github.com/fire833/vroute/pkg/client/clientset"
	certificates "github.com/fire833/vroute/pkg/client/listers/apis/certificates"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CertificateSignerListInformer provides access to a shared informer and lister for
// CertificateSignerLists.
type CertificateSignerListInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() certificates.CertificateSignerListLister
}

type certificateSignerListInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCertificateSignerListInformer constructs a new informer for CertificateSignerList type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCertificateSignerListInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCertificateSignerListInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCertificateSignerListInformer constructs a new informer for CertificateSignerList type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCertificateSignerListInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesCertificates().CertificateSignerLists(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesCertificates().CertificateSignerLists(namespace).Watch(context.TODO(), options)
			},
		},
		&apiscertificates.CertificateSignerList{},
		resyncPeriod,
		indexers,
	)
}

func (f *certificateSignerListInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCertificateSignerListInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *certificateSignerListInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiscertificates.CertificateSignerList{}, f.defaultInformer)
}

func (f *certificateSignerListInformer) Lister() certificates.CertificateSignerListLister {
	return certificates.NewCertificateSignerListLister(f.Informer().GetIndexer())
}
