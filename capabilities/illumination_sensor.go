package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
)

// IlluminationReading is an Illumination reading.
type IlluminationReading struct {
	// Value contains a floating point number representing the Illumination, in lux.
	Value float64
}

// IlluminationSensor is a capability that provides Illumination readings from a device. Multiple readings can be returned
// if the device has multiple sensors.
type IlluminationSensor interface {
	// Reading reads (or provides the most recent) Illumination readings the device has.
	Reading(context.Context) ([]IlluminationReading, error)
}

// IlluminationSensorState is sent to inform consumers of the devices Illumination values, there may be no change.
type IlluminationSensorState struct {
	// Device that is informing of its Illumination readings state.
	Device da.Device
	// New state of device.
	State []IlluminationReading
}
