package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
)

// RelativeHumidityReading is a relative humidity reading.
type RelativeHumidityReading struct {
	// Value contains a floating point number representing the relative humidity, between 0.0 and 1.0.
	Value float64
}

// RelativeHumiditySensor is a capability that provides the relative humidity readings from a device. Multiple readings
// can be returned if the device has multiple sensors.
type RelativeHumiditySensor interface {
	// Reading reads (or provides the most recent) relative humidity readings the device has.
	Reading(context.Context) ([]RelativeHumidityReading, error)
}

// RelativeHumiditySensorUpdate is sent to inform consumers of the devices humidity values, there may be no change.
type RelativeHumiditySensorUpdate struct {
	// Device that is informing of its humidity readings state.
	Device da.Device
	// New state of device.
	State []RelativeHumidityReading
}
