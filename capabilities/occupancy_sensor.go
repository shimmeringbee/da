package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
)

// OccupancyReading is a Occupancy reading.
type OccupancyReading struct {
	// Occupied is true if the sensor detects the space the sensor serves is occupied.
	Occupied bool
}

// OccupancySensor is a capability that provides Occupancy readings from a device. Multiple readings can be returned
// if the device has multiple sensors.
type OccupancySensor interface {
	// Reading reads (or provides the most recent) Occupancy readings the device has.
	Reading(context.Context, da.Device) ([]OccupancyReading, error)
}

// OccupancySensorState is sent to inform consumers of the devices Occupancy values, there may be no change.
type OccupancySensorState struct {
	// Device that is informing of its Occupancy readings state.
	Device da.Device
	// New state of device.
	State []OccupancyReading
}
