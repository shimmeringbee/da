package capabilities

import (
	"context"
	. "github.com/shimmeringbee/da"
)

// EnumerateCapabilities is a capability which allows the user of an abstraction to request that a devices
// capabilities be enumerated, this may be called on a device which has previously been enumerated.
//
// Gateways will send EnumerateCapabilitiesComplete messages to announce a devices capabilities regardless
// of having the EnumerateCapabilities capability.
type EnumerateCapabilities interface {
	// Enumerate forces the gateway to rescan the device specified and enumerate its capabilities.
	Enumerate(context.Context, Device) error
}

// Event sent to inform consumers that capability enumeration is beginning on a device.
type EnumerateCapabilitiesBegin struct {
	// Device that enumeration failed on.
	Device
}

// Event sent to inform consumers that capability enumeration failed on a device.
type EnumerateCapabilitiesFailed struct {
	// Device that enumeration failed on.
	Device
}

// Event sent to inform consumers that capability enumeration has completed on a device.
type EnumerateCapabilitiesComplete struct {
	// Device enumeration completed.
	Device
}
