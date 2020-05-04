package capabilities

import (
	. "github.com/shimmeringbee/da"
	"time"
)

type DeviceDiscovery interface {
	Allow(Device, time.Duration) error
	Deny(Device) error
	Status(Device) (bool, time.Duration)
}

type DeviceDiscoveryAllowed struct {
	Device
	Duration time.Duration
}

type DeviceDiscoveryDenied struct {
	Device
}
