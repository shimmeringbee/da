package capabilities

import (
	"context"
	. "github.com/shimmeringbee/da"
	"time"
)

type DeviceDiscovery interface {
	Allow(context.Context, Device, time.Duration) error
	Deny(context.Context, Device) error
	Status(context.Context, Device) (DeviceDiscoveryStatus, error)
}

type DeviceDiscoveryStatus struct {
	Discovering       bool
	RemainingDuration time.Duration
}

type DeviceDiscoveryAllowed struct {
	Device
	Duration time.Duration
}

type DeviceDiscoveryDenied struct {
	Device
}
