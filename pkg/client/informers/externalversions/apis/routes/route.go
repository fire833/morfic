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

package routes

import (
	"context"
	time "time"

	apisroutes "github.com/fire833/morfic/pkg/apis/routes"
	versioned "github.com/fire833/morfic/pkg/client/clientset/versioned"
	internalinterfaces "github.com/fire833/morfic/pkg/client/informers/externalversions/internalinterfaces"
	routes "github.com/fire833/morfic/pkg/client/listers/apis/routes"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RouteInformer provides access to a shared informer and lister for
// Routes.
type RouteInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() routes.RouteLister
}

type routeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRouteInformer constructs a new informer for Route type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRouteInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRouteInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRouteInformer constructs a new informer for Route type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRouteInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RoutesRoutes().Routes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RoutesRoutes().Routes(namespace).Watch(context.TODO(), options)
			},
		},
		&apisroutes.Route{},
		resyncPeriod,
		indexers,
	)
}

func (f *routeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRouteInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *routeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisroutes.Route{}, f.defaultInformer)
}

func (f *routeInformer) Lister() routes.RouteLister {
	return routes.NewRouteLister(f.Informer().GetIndexer())
}
