package capabilities

import (
	"context"
	. "github.com/shimmeringbee/da"
)

// PressureReading is an pressure reading.
type PressureReading struct {
	// Value contains a floating point number representing the pressure, in pascal.
	Value float32
}

// PressureSensor is a capability that provides pressure readings from a device. Multiple readings can be returned if
// the device has multiple sensors.
type PressureSensor interface {
	// Reading reads (or provides the most recent) pressure readings the device has.
	Reading(context.Context, Device) ([]PressureReading, error)
}
