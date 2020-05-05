package capabilities

import (
	"context"
	. "github.com/shimmeringbee/da"
	"time"
)

type DeviceDiscovery interface {
	Enable(context.Context, Device, time.Duration) error
	Disable(context.Context, Device) error
	Status(context.Context, Device) (DeviceDiscoveryStatus, error)
}

type DeviceDiscoveryStatus struct {
	Discovering       bool
	RemainingDuration time.Duration
}

type DeviceDiscoveryAllowed struct {
	Gateway
	Duration time.Duration
}

type DeviceDiscoveryDenied struct {
	Gateway
}
