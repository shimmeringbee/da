package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
)

// EnumerateDevice is a capability which allows the user of an abstraction to request that a devices
// capabilities and facts be enumerated, this may be called on a device which has previously been
// enumerated.
//
// Gateways will send EnumerateDeviceSuccess messages to announce a devices capabilities regardless
// of having the EnumerateDevice capability.
type EnumerateDevice interface {
	// Enumerate forces the gateway to rescan the device specified and enumerate its capabilities.
	//
	// This must not be blocking, errors which occur during enumeration should cause EnumerateDeviceFailure
	// to be sent, if enumeration could not be recovered.
	Enumerate(context.Context, da.Device) error

	// Retrieve the current enumeration status for the device.
	Status(context.Context, da.Device) (EnumerationStatus, error)
}

// Status of devices enumeration.
type EnumerationStatus struct {
	// True if device is being enumerated.
	Enumerating bool
}

// Event sent to inform consumers that enumeration is beginning on a device.
type EnumerateDeviceStart struct {
	// Device that enumeration failed on.
	da.Device
}

// Event sent to inform consumers that enumeration failed on a device.
type EnumerateDeviceFailure struct {
	// Device that enumeration failed on.
	da.Device
	// Error as to why enumeration failed.
	Error error
}

// Event sent to inform consumers that enumeration has completed on a device.
type EnumerateDeviceSuccess struct {
	// Device enumeration completed.
	da.Device
}
