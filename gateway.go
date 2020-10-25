package da

import "context"

// A gateway is a provider of devices and Capability objects to control them
type Gateway interface {
	// Read events from the gateway about devices, this must always be serviced. context.DeadlineExceeded will be
	// returned if the context expires during a read. Use this mechanism if you can not indefinitely wait for an event.
	ReadEvent(context.Context) (interface{}, error)

	// Retrieve a capability implementation from the gateway, may return nil if capability is not supported.
	Capability(Capability) interface{}

	// Fetch a device which represents the gateway its self, this allows a gateway to have optional capabilities.
	Self() Device

	// Fetch a list of all devices on the gateway, including its self.
	Devices() []Device

	// Start the gateway, including any background routines required.
	Start() error

	// Stop the gateway, including any background routines required. Gateways are not required to be reusable,
	// and a new gateway should be created if a restart is required.
	Stop() error
}
