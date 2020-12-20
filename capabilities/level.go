package capabilities

import (
	"context"
	. "github.com/shimmeringbee/da"
	"time"
)

// Level is a capability that provides the ability for a device to set the level of an output, this
// could be brightness, fan intensity, blind height, desk height, etc.
type Level interface {
	// Change changes the devices current output level to the new required value, over a set duration.
	Change(context.Context, Device, float64, time.Duration) error
	// Status returns the current status of Level.
	Status(context.Context, Device) (LevelStatus, error)
}

// LevelStatus represents the current level.
type LevelStatus struct {
	// CurrentLevel is the current level of the device.
	CurrentLevel float64
	// TargetLevel is the target level of the device.
	TargetLevel float64
	// DurationRemaining is how long until CurrentLevel transitions to TargetLevel.
	DurationRemaining time.Duration
}

// LevelStatusUpdate is sent to inform consumers of the devices level state, may be sent even if there is no change.
type LevelStatusUpdate struct {
	// Device that is informing of it's on off state.
	Device Device
	// New state of device.
	State LevelStatus
}
