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
	"github.com/fire833/vroute/pkg/apis/interfaces"
	"github.com/fire833/vroute/pkg/controller"
)

const (
	chanQ int = 3

	numberOfWorkers
)

type LinkListChan chan interfaces.LinkList
type LinkChan chan interfaces.Link

type InterfaceController struct {
	linkListChannel LinkListChan
	linkChannel     LinkChan

	// StopChan is the channel that should send exit codes to threads to close out after
	// they have completed their current work.
	stopChan controller.StopChan
}

func (c *InterfaceController) BeginWorkers(num int) {
	lchan := make(LinkChan, chanQ)
	llchan := make(LinkListChan, chanQ)

	for w := 0; w < num; w++ {
		go c.beginInterfaceControllerWorker(lchan, llchan)
	}
}

func (c *InterfaceController) GracefulStop(err uint8) {
	c.stopChan <- err
}

func (c *InterfaceController) RunWorkers() int {
	return numberOfWorkers
}

func (c *InterfaceController) beginInterfaceControllerWorker(link LinkChan, list LinkListChan) {
	for {
		select {
		case linkevent := <-link:
			{

			}
		case linklistevent := <-list:
			{

			}
		case <-c.stopChan:
			{
				return // Just return for now since we need to kill the thread.
			}
		}
	}
}
