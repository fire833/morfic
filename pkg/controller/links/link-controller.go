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

package interfaces

import (
	"context"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"

	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type LinkController struct {
	client        clientset.Interface
	eventRecorder record.EventRecorder

	lListerSynched cache.InformerSynced

	// Deployments that need to be synced
	queue workqueue.RateLimitingInterface

	// workerLoopPeriod is the time between worker runs. The workers process the queue of service and pod changes.
	workerLoopPeriod time.Duration
}

// NewLinkConntroller creates a new LinkController
func NewLinkController(client clientset.Interface) (*LinkController, error) {
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: client.CoreV1().Events("")}) // Need to implement local clientset.

	lc := &LinkController{
		client:        client,
		eventRecorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "link-controller"}),
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "link"),
	}

	return lc, nil

}

// Run starts off the number of worker threads as defined by workers.
// Each thread watches and syncs state of Link objects.
func (c *LinkController) Run(ctx context.Context, workers int) {
	defer c.queue.ShutDown()

	for i := 0; i < workers; i++ {
		go wait.UntilWithContext(ctx, c.worker, c.workerLoopPeriod)
	}

	<-ctx.Done()
}

// Pops events off of the workqueues, processes them, and then marks
// them as done. Is thread safe due tothe worker queue implementation.
func (c *LinkController) worker(ctx context.Context) {
	for c.processNextWorkItem(ctx) {
	}
}

func (c *LinkController) processNextWorkItem(ctx context.Context) bool {
	lkey, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(lkey)
	c.syncLink(ctx, lkey.(string))
	c.queue.Forget(lkey)

	return true
}

func (c *LinkController) addLink() {

}

func (c *LinkController) delLink() {

}

func (c *LinkController) updateLink() {

}

func (c *LinkController) syncLink(ctx context.Context, key string) error {
	return nil
}
