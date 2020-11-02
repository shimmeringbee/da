package capabilities

import (
	"context"
	. "github.com/shimmeringbee/da"
)

// OnOff is a capability which signifies that a device has the ability to be switched on or off.
type OnOff interface {
	// On turns the device on, no operation if device is already on.
	On(context.Context, Device) error

	// Off turns the device off, no operation if device is already off.
	Off(context.Context, Device) error

	// State returns the current device state.
	Status(context.Context, Device) (bool, error)
}

// OnOffState is sent to inform consumers of the devices on/off state, may be sent even if there is no change.
type OnOffState struct {
	// Device that is informing of it's on off state.
	Device Device
	// New state of device.
	State bool
}
