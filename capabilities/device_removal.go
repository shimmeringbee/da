package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
)

// RemovalType enumerates different possibilities for modes of removal.
type RemovalType uint8

const (
	// Request removals may not succeed, the device may need to be cooperative or online.
	Request RemovalType = iota
	// Force removals must occur, this may result in a dirty state, such as Zigbee devices being able to rejoin
	// as they have not wiped their key.
	Force
)

// DeviceRemoval provides a capability for a device to be removed from a gateway upon request.
type DeviceRemoval interface {
	// Remove requests that the device is removed from the gateway.
	Remove(context.Context, da.Device, RemovalType) error
}
