package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
)

// PowerSupply exposes how a device is powered.
type PowerSupply interface {
	// Status provides the devices current state of the devices power supply
	Status(context.Context) (PowerState, error)
}

// PowerState represents the state of all power sources for the device.
type PowerState struct {
	// Mains represents any mains powered supplies, there may be multiple if a device has A/B supplies.
	Mains []PowerMainsState
	// Battery represents any battery supplies, there may be multiple.
	Battery []PowerBatteryState
}

type PowerStatusPresent uint

const (
	Voltage PowerStatusPresent = 1 << iota
	Frequency
	MaximumVoltage
	MinimumVoltage
	Remaining
	Available
)

// PowerMainsState describes a mains power supply for the device.
type PowerMainsState struct {
	// Voltage of mains supply.
	Voltage float64
	// Frequency of mains supply, 0 if DC.
	Frequency float64
	// Available represents if the supply is valid and usable.
	Available bool

	// Present represents a bitmap of the fields present.
	Present PowerStatusPresent
}

// PowerBatteryState describes a battery power supply for the device.
type PowerBatteryState struct {
	// Voltage of battery.
	Voltage float64
	// MaximumVoltage of battery.
	MaximumVoltage float64
	// MinimumVoltage of battery, i.e. empty or unable to support device.
	MinimumVoltage float64

	// Remaining ratio of battery, 0.0 at empty, 1.0 at full charge.
	Remaining float64
	// Available represents if the supply is valid and usable.
	Available bool

	// Present represents a bitmap of the fields present.
	Present PowerStatusPresent
}

// PowerStatusUpdate is sent to inform consumers that an update to the power supply state of the device has changed.
type PowerStatusUpdate struct {
	// Device which has an update.
	Device da.Device
	// PowerStatus of device.
	PowerStatus PowerState
}
