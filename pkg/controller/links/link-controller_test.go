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

package interfaces

import (
	"context"
	"reflect"
	"testing"
	"time"

	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
)

func TestNewLinkController(t *testing.T) {
	type args struct {
		client clientset.Interface
	}
	tests := []struct {
		name    string
		args    args
		want    *LinkController
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLinkController(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLinkController() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkController_processNextWorkItem(t *testing.T) {
	type fields struct {
		client           clientset.Interface
		eventRecorder    record.EventRecorder
		lListerSynched   cache.InformerSynced
		queue            workqueue.RateLimitingInterface
		workerLoopPeriod time.Duration
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LinkController{
				client:           tt.fields.client,
				eventRecorder:    tt.fields.eventRecorder,
				lListerSynched:   tt.fields.lListerSynched,
				queue:            tt.fields.queue,
				workerLoopPeriod: tt.fields.workerLoopPeriod,
			}
			if got := c.processNextWorkItem(tt.args.ctx); got != tt.want {
				t.Errorf("LinkController.processNextWorkItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
