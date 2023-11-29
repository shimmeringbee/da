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
	// Enumerate forces the gateway to rescan the device and enumerate its capabilities.
	//
	// This must not be blocking, errors which occur during enumeration should cause EnumerateDeviceFailure
	// to be sent, if enumeration could not be recovered.
	Enumerate(context.Context) error

	// Status retrieves the current enumeration status for the device.
	Status(context.Context) (EnumerationStatus, error)
}

type EnumerationCapability struct {
	Attached bool
	Error    error
}

// EnumerationStatus of devices enumeration.
type EnumerationStatus struct {
	// Enumerating represents if device is being enumerated.
	Enumerating bool
	// CapabilityStatus shows if a capability successfully attached, and/or if an error was raised.
	CapabilityStatus map[da.Capability]EnumerationCapability
}

// EnumerateDeviceStart sent to inform consumers that enumeration is beginning on a device.
type EnumerateDeviceStart struct {
	// Device that enumeration failed on.
	Device da.Device
}

// EnumerateDeviceStopped sent to inform consumers that enumeration has completed on a device.
type EnumerateDeviceStopped struct {
	// Device enumeration completed.
	Device da.Device
	// Status is the result of the enumeration.
	Status EnumerationStatus
}
