package capabilities

import (
	"context"
	. "github.com/shimmeringbee/da"
)

// DeviceRemoval provides a capability for a device to be removed from a gateway upon request.
type DeviceRemoval interface {
	// Remove requests that the device is removed from the gateway.
	Remove(context.Context, Device) error
}
