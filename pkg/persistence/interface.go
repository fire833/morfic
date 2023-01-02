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

package persistence

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

/*
	Persistence provider is an interface for storage backends to persist data
	about the control plane to a durable medium for persistence across restarts
	of both the service and the bare metal as a whole. Any persistence provider
	must be able to fulfill the methods described in this interface, otherwise
	the control plane will reject the provider and not be able to communicate/use
	the interface.

	Data written/retrieved from this provider should be considered generally not very
	sensitive, but should still be protected from being exposed to the public/
	over the network. Thus, this interface will be used for storing general
	configuration data/metadata, rather than managing secrets such as API keys
	and credentials. Storage of sensitive information such as the ones perviously
	described will be set to be managed by the SecretPersistenceProviders that are
	registered/configured with the control plane.
*/
type PersistenceProvider interface {
	/*
		Initialize() is designed to be called by the control plane to initialize
		any connections/files or any other initialization that needs to be
		completed by the individual persistence provider to get ready for serving
		requests from the control plane.

		Examples for tasks to be implemented by a provider during Initialize should
		include:

		* Opening/reading file that was persisted and parsing into memory for immediate
		  consumption by the control plane.
		* Creating a client connection to the remote endpoint that will pull the data
		  (ie starting an s3.S3 connection locally, etc.).

		If there is an error in initializing this backend, the appropriate error will
		be returned and the control plane will assume the backend is dead.
	*/
	Initialize(id *MachineId) error

	/*
		Instances returns to the controlplane the names of the instances this
		PersistenceProvider is currently supporting. This is especially useful
		for having redundant backends that are manager by the same provider.
		An example being you want to persist storage to two separate S3 APIs,
		AWS S3, as well as a local MinIO instance hosted by your organization.
		Another example could be you want to persist to both a file locally mounted
		on the bare metal machine, and in another location of the filesystem that is
		actually mounted via NFS onto the bare metal host.

		It is generally up to the discretion of the provider to determine how
		individual instances are to be labeled, but names should be unique and avoid
		any chances for collision between two or more instances at runtime. Given that
		these strings are only to be parsed by persistence management layers rather than
		human administrators, so they do not have to be human readable. The suggested format
		is to repesent an instance as <generic_instance_name>-<instance_hash>.
		To copy the s3 instance noted above, generic_instance_name could be s3, and the
		instance_hash be a uniquely generated hex hash from the configuration of that
		instance.
	*/
	Instances() []string

	/*
		InstanceStatus allows the control plane to asyncronously access the status for
		a specific known instance of a backend. This can be useful for when the control
		plane knows it will need to make a bunch of write/read requests (especially
		on startup/shutdown) and wants to have more confidence that those instances are
		up before performing operations on them.
	*/
	InstanceStatus(instance string) *PersistenceProviderStatus

	/*
		PutObject persists a runtime object to this persistence provider. Callers
		should specify the specific instance to which this provider should persist the object,
		and the object itself, which conforms to the runtime.Object interface.

		This method returns the key to accessthis specific object in the future, so it can
		be cached by the persistence management layer with minimal memory footprint. If there
		was an error in persisting the data, a corresponding error shall be returned.
	*/
	PutObject(instance string, object *runtime.Object) (key ObjectKey, e error)

	/*
		DeleteObject removes a runtime object from a specific provider instance.
	*/
	DeleteObject(instance string, kind schema.GroupVersionKind, key ObjectKey) error

	/*
		GetObject returns a persisted instance of an object from a specific provider instance.
	*/
	GetObject(instance string, kind schema.GroupVersionKind, key ObjectKey) *runtime.Object

	/*
		GetAllObjects returns all of the persisted instances of this object kind stored with
		this persistence provider. The ObjectKind denotes the types that callers request, and
		the instance should be a valid provider instance to be queried for the data.
	*/
	GetAllObjects(instance string, kind schema.GroupVersionKind) []*runtime.Object
}

/*
	PersistenceProviderStatus is a structure that is returned from a PersistenceProvider
	after the controlplane asks for the status of a particular instance under the management
	domain for a provider.
*/
type PersistenceProviderStatus struct {

	// Returns true if the string instance exists. Otherwise returns false. The control plane
	// should assume that this instance doesn't exist if this field is set to false and abort
	// any planned upcoming operations to the instance.
	InstanceExists bool `json:"exists" yaml:"exists"`

	// Returns true if the instance is reported as up by the provider. Otherwise should be
	// considered down and the control plane should abort any planned upcoming operations
	// on the instance.
	Status bool `json:"status,omitempty" yaml:"status,omitempty"`

	// Returns the average persistence persistence latency for this instance. This latency should
	// be measured as the time from when the request for an object to persist gets called, and
	// the time where the persistence backend has processed the request (ie s3 has returned a 200
	// or 201 from a post/delete request, file has been been saved, basically transaction has concluded).
	MeanPersistLatency float64 `json:"mean_persist_latency,omitempty" yaml:"meanPersistLatency,omitempty"`
}

/*
	SecretPersistenceProvider serves as an interface for storing secrets from the control plane
	to a persistent (and secure) medium.
*/
type SecretPersistenceProvider interface {
}
