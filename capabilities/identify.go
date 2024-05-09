package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
	"time"
)

// Identify is a capability which signifies that a device has the ability to identify itself to the end user.
type Identify interface {
	// Identify causes the device to identify itself to the end user.
	Identify(context.Context, time.Duration) error

	// Status returns the current device state.
	Status(context.Context) (bool, error)
}

// IdentifyState is sent to inform consumers of the devices identify state, may be sent even if there is no change.
type IdentifyState struct {
	// Device that is informing of it's on off state.
	Device da.Device
	// New state of device.
	State bool
}
