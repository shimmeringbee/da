package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
)

// TemperatureReading is a temperature reading.
type TemperatureReading struct {
	// Value contains a floating point number representing the temperature, in Kelvin.
	Value float64
}

// TemperatureSensor is a capability that provides temperature readings from a device. Multiple readings can be returned
// if the device has multiple sensors.
type TemperatureSensor interface {
	// Reading reads (or provides the most recent) temperature readings the device has.
	Reading(context.Context) ([]TemperatureReading, error)
}

// TemperatureSensorState is sent to inform consumers of the devices temperature values, there may be no change.
type TemperatureSensorState struct {
	// Device that is informing of its temperature readings state.
	Device da.Device
	// New state of device.
	State []TemperatureReading
}
