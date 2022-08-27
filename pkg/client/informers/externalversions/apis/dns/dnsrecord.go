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

package dns

import (
	"context"
	time "time"

	apisdns "github.com/fire833/morfic/pkg/apis/dns"
	versioned "github.com/fire833/morfic/pkg/client/clientset/versioned"
	internalinterfaces "github.com/fire833/morfic/pkg/client/informers/externalversions/internalinterfaces"
	dns "github.com/fire833/morfic/pkg/client/listers/apis/dns"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DNSRecordInformer provides access to a shared informer and lister for
// DNSRecords.
type DNSRecordInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() dns.DNSRecordLister
}

type dNSRecordInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDNSRecordInformer constructs a new informer for DNSRecord type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDNSRecordInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDNSRecordInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDNSRecordInformer constructs a new informer for DNSRecord type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDNSRecordInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DnsV1alpha1().DNSRecords(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DnsV1alpha1().DNSRecords(namespace).Watch(context.TODO(), options)
			},
		},
		&apisdns.DNSRecord{},
		resyncPeriod,
		indexers,
	)
}

func (f *dNSRecordInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDNSRecordInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dNSRecordInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisdns.DNSRecord{}, f.defaultInformer)
}

func (f *dNSRecordInformer) Lister() dns.DNSRecordLister {
	return dns.NewDNSRecordLister(f.Informer().GetIndexer())
}
