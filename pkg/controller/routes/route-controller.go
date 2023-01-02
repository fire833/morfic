/*
*	Copyright (C) 2023 Kendall Tauser
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

package routes

import (
	"time"

	clientset "github.com/fire833/morfic/pkg/client/clientset/versioned"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type RouteController struct {
	client clientset.Interface

	// queue of new work items to be processed.
	queue workqueue.RateLimitingInterface

	// routeLister

	routeSynched cache.InformerSynced
}

func (c *RouteController) Run(threadiness int, stopCh chan struct{}) {

	// Defer shutting down the controller workqueue.
	defer c.queue.ShutDown()

	for i := 0; i < threadiness; i++ {
		// Start each worker on a new goroutine.
		go wait.Until(c.routeWorker, time.Second, stopCh)
	}

	// Block until we want to stop the controller.
	<-stopCh
}

func (c *RouteController) routeWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *RouteController) processNextWorkItem() bool {

	key, quit := c.queue.Get()
	if quit {
		return false
	}

	// Defer finishing the processing of this key from the queue.
	defer c.queue.Done(key)

	return false
}
