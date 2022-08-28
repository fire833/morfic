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

package net

import (
	"context"
	internalinterfaces "pkg/client/informers/externalversions/internalinterfaces"
	time "time"

	apisnet "github.com/fire833/morfic/pkg/apis/net"
	versioned "github.com/fire833/morfic/pkg/client/clientset/versioned"
	net "github.com/fire833/morfic/pkg/client/listers/apis/net"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RouteTableInformer provides access to a shared informer and lister for
// RouteTables.
type RouteTableInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() net.RouteTableLister
}

type routeTableInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewRouteTableInformer constructs a new informer for RouteTable type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRouteTableInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRouteTableInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredRouteTableInformer constructs a new informer for RouteTable type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRouteTableInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetNet().RouteTables().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetNet().RouteTables().Watch(context.TODO(), options)
			},
		},
		&apisnet.RouteTable{},
		resyncPeriod,
		indexers,
	)
}

func (f *routeTableInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRouteTableInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *routeTableInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisnet.RouteTable{}, f.defaultInformer)
}

func (f *routeTableInformer) Lister() net.RouteTableLister {
	return net.NewRouteTableLister(f.Informer().GetIndexer())
}
