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

import "github.com/fire833/vroute/pkg/apis/interfaces"

const chanQ int = 3

type LinkListChan chan interfaces.LinkList
type LinkChan chan interfaces.Link

type InterfaceController struct {
	linkListChannels map[int]LinkListChan
	linkChannels     map[int]LinkChan
}

func (c *InterfaceController) BeginWorkers(num int) {
	for w := 0; w < num; w++ {
		lchan := make(LinkChan, chanQ)
		llchan := make(LinkListChan, chanQ)

		c.linkChannels[w] = lchan
		c.linkListChannels[w] = llchan

		go c.beginInterfaceControllerWorker(lchan, llchan)
	}
}

func (c *InterfaceController) GracefulStop() error {
	return nil
}

func (c *InterfaceController) ForceStop() error {
	return nil
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
		}
	}
}
