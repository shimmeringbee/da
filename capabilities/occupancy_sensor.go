package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
	"time"
)

// OccupancyReading is an Occupancy reading.
type OccupancyReading struct {
	// Occupied is true if the sensor detects the space the sensor serves is occupied.
	Occupied bool
	// Duration is the length of time the occupancy has been detected.
	Duration time.Duration
}

// OccupancySensor is a capability that provides Occupancy readings from a device. Multiple readings can be returned
// if the device has multiple sensors.
type OccupancySensor interface {
	// Reading reads (or provides the most recent) Occupancy readings the device has.
	Reading(context.Context) ([]OccupancyReading, error)
}

// OccupancySensorUpdate is sent to inform consumers of the devices Occupancy values, there may be no change.
type OccupancySensorUpdate struct {
	// Device that is informing of its Occupancy readings state.
	Device da.Device
	// New state of device.
	State []OccupancyReading
}
