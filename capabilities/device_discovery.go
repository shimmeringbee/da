package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
	"time"
)

// DeviceDiscovery is a capability which allows the user of the gateway to control if new devices
// are discovered. This may be an active process when the gateway seeks out devices, or passive
// where the gateway now permits joining of new devices.
//
// The absence of this capability does not prevent a gateway autonomously introducing a new device.
type DeviceDiscovery interface {
	// Enable device discovery on the gateway for the duration specified, after the duration
	// has elapsed device discovery will be disabled.
	//
	// If Enable is called on a gateway which is already enabled, then the timeout will be
	// reset to the new duration from that point in time, not the original enablement time.
	Enable(context.Context, da.Device, time.Duration) error

	// Disable device discovery on the gateway.
	Disable(context.Context, da.Device) error

	// Retrieve the current discovery status for the gateway.
	Status(context.Context, da.Device) (DeviceDiscoveryStatus, error)
}

// Status of gateways device discovery.
type DeviceDiscoveryStatus struct {
	// True if gateway is discovering.
	Discovering bool

	// Remaining duration of discovery, will be 0 if not discovering.
	RemainingDuration time.Duration
}

// Event sent to inform abstraction consumer that the gateway is discovering new devices, this may be sent multiple
// times without a DeviceDiscoveryDisabled between if the discovery duration is adjusted.
type DeviceDiscoveryEnabled struct {
	// Gateway discovery is enabled on.
	da.Gateway

	// Duration that the discovery will be performed for.
	Duration time.Duration
}

// Event sent to inform abstraction consumer that the gateway is no longer discovering.
type DeviceDiscoveryDisabled struct {
	// Gateway discovery is disabled on.
	da.Gateway
}
