package da

import (
	"context"
)

// Gateway is a provider of devices and Capability objects to control them
type Gateway interface {
	// ReadEvent from the gateway about devices, this must always be serviced. context.DeadlineExceeded will be
	// returned if the context expires during a read. Use this mechanism if you can not indefinitely wait for an event.
	ReadEvent(context.Context) (interface{}, error)

	// Capability gets a capability implementation from the gateway, may return nil if capability is not supported.
	Capability(Capability) interface{}

	// Capabilities returns a list of all Capabilities that the gateway can support.
	Capabilities() []Capability

	// Self fetches a device which represents the gateway, this allows a gateway to have optional capabilities.
	Self() Device

	// Devices fetch a list of all devices on the gateway, including its self.
	Devices() []Device

	// Start the gateway, including any background routines required. The context is only for the process of starting
	// a gateway.
	Start(context.Context) error

	// Stop the gateway, including any background routines required. Gateways are not required to be reusable,
	// and a new gateway should be created if a restart is required. The context is only for the process of stopping
	// a gateway.
	Stop(context.Context) error
}
