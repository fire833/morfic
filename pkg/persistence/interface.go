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

package persistence

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
	Initialize() error

	/*
		Instances returns to the controlplane the names of the instances this
		PersistenceProvider is currently supporting. This is especially useful
		for having redundant backends that are manager by the same provider.
		An example being you want to persist storage to two separate S3 APIs,
		AWS S3, as well as a local MinIO instance hosted by your organization.
		Another example could be you want to persist to both a file locally mounted
		on the bare metal machine, and in another location of the filesystem that is
		actually mounted via NFS onto the bare metal host.
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

	// WORK IN PROGRESS

	GetObjects(instance string)

	PutObjects(instance string)

	DeleteObjects(instance string)

	UpdateObjects(instance string)
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
}

/*

 */
type SecretPersistenceProvider interface {
}
