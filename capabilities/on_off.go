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

	// Toggle changes the devices state to be opposite, off -> on, on -> off.
	Toggle(context.Context, Device) error

	// State returns the current device state.
	State(context.Context, Device) (bool, error)
}

// OnOffState is sent to inform consumers of the devices on/off state, may be sent even if there is no change.
type OnOffState struct {
	// Device that enumeration failed on.
	Device Device
	// New state of device.
	State bool
}
