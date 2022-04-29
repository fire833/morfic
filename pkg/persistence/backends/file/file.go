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

package file

import (
	"crypto/rand"
	"fmt"

	"github.com/fire833/vroute/pkg/persistence"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// FilePersistenceProvider represents a persistence provider that is able to
// persist control-plane data to a file, and can retrieve data according to the
// PersistenceProvider interface.
type FilePersistenceProvider struct {
	id           *persistence.MachineId
	instances    []string
	writeLatency float64
	up           bool
}

func (file *FilePersistenceProvider) Initialize(id *persistence.MachineId) {
	file.id = id
	file.up = true
	file.writeLatency = 0

	// Start off with only one provider available for now.
	file.instances = make([]string, 1)

	buf := make([]byte, 10)
	rand.Read(buf)

	file.instances[1] = fmt.Sprintf("file-%b", buf)
}

func (file *FilePersistenceProvider) Instances() []string {
	return file.instances
}

func (file *FilePersistenceProvider) InstanceStatus(instance string) *persistence.PersistenceProviderStatus {
	for n, instance := range file.instances {
		if instance == file.instances[n] {
			return &persistence.PersistenceProviderStatus{
				InstanceExists:     true,
				Status:             file.up,
				MeanPersistLatency: file.writeLatency,
			}
		}
	}

	return &persistence.PersistenceProviderStatus{
		InstanceExists:     false,
		Status:             false,
		MeanPersistLatency: 0,
	}
}

func (file *FilePersistenceProvider) PutObject(instance string, object *runtime.Object) (key persistence.ObjectKey, e error) {
	return nil, nil
}

func (file *FilePersistenceProvider) DeleteObject(instance string, kind schema.GroupVersionKind, key persistence.ObjectKey) error {
	return nil
}

func (file *FilePersistenceProvider) GetObject(instance string, kind schema.GroupVersionKind, key persistence.ObjectKey) *runtime.Object {
	return nil
}

func (file *FilePersistenceProvider) GetAllObjects(instance string, kind schema.GroupVersionKind) []*runtime.Object {
	return nil
}
