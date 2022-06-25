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

package v1

import (
	"context"
	internalinterfaces "github.com/fire833/morfic/pkg/client/informers/internalinterfaces"
	time "time"

	apisinterfaces "github.com/fire833/morfic/pkg/apis/interfaces"
	clientset "github.com/fire833/morfic/pkg/client/clientset"
	interfaces "github.com/fire833/morfic/pkg/client/listers/interfaces"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LinkInformer provides access to a shared informer and lister for
// Links.
type LinkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() interfaces.LinkLister
}

type linkInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLinkInformer constructs a new informer for Link type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLinkInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLinkInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLinkInformer constructs a new informer for Link type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLinkInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InterfacesInterfaces().Links(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InterfacesInterfaces().Links(namespace).Watch(context.TODO(), options)
			},
		},
		&apisinterfaces.Link{},
		resyncPeriod,
		indexers,
	)
}

func (f *linkInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLinkInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *linkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisinterfaces.Link{}, f.defaultInformer)
}

func (f *linkInformer) Lister() interfaces.LinkLister {
	return interfaces.NewLinkLister(f.Informer().GetIndexer())
}
