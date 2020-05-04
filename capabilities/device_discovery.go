package capabilities

import (
	"time"
)

type DeviceDiscovery interface {
	Allow(time.Duration)
	Deny()
	Status() (bool, time.Duration)
}
