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
	internalinterfaces "pkg/client/informers/externalversions/internalinterfaces"
	time "time"

	apisdns "github.com/fire833/morfic/pkg/apis/dns"
	versioned "github.com/fire833/morfic/pkg/client/clientset/versioned"
	dns "github.com/fire833/morfic/pkg/client/listers/apis/dns"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DNSRecordListInformer provides access to a shared informer and lister for
// DNSRecordLists.
type DNSRecordListInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() dns.DNSRecordListLister
}

type dNSRecordListInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDNSRecordListInformer constructs a new informer for DNSRecordList type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDNSRecordListInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDNSRecordListInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDNSRecordListInformer constructs a new informer for DNSRecordList type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDNSRecordListInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DnsDns().DNSRecordLists().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DnsDns().DNSRecordLists().Watch(context.TODO(), options)
			},
		},
		&apisdns.DNSRecordList{},
		resyncPeriod,
		indexers,
	)
}

func (f *dNSRecordListInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDNSRecordListInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dNSRecordListInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisdns.DNSRecordList{}, f.defaultInformer)
}

func (f *dNSRecordListInformer) Lister() dns.DNSRecordListLister {
	return dns.NewDNSRecordListLister(f.Informer().GetIndexer())
}
