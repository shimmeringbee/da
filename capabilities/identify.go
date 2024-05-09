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

// IdentifyState contains information regarding the state of identify.
type IdentifyState struct {
	// Identifying is true if the device is currently identifying.
	Identifying bool
	// Remaining is the duration will continue for if Identifying is true.
	Remaining time.Duration
}

// IdentifyUpdate is sent to inform consumers of the devices identify state, may be sent even if there is no change.
type IdentifyUpdate struct {
	// Device that is informing of it's on off state.
	Device da.Device
	// State of device.
	State IdentifyState
}
