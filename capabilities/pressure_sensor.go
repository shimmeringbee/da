package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
)

// PressureReading is a pressure reading.
type PressureReading struct {
	// Value contains a floating point number representing the pressure, in pascal.
	Value float64
}

// PressureSensor is a capability that provides pressure readings from a device. Multiple readings can be returned if
// the device has multiple sensors.
type PressureSensor interface {
	// Reading reads (or provides the most recent) pressure readings the device has.
	Reading(context.Context) ([]PressureReading, error)
}

// PressureSensorUpdate is sent to inform consumers of the devices pressure values, there may be no change.
type PressureSensorUpdate struct {
	// Device that is informing of its pressure readings state.
	Device da.Device
	// New state of device.
	State []PressureReading
}
