package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
	"time"
)

// AlertType represents a type of alert a warning device can make.
type AlertType uint8

func (a AlertType) String() string {
	if name, found := AlertTypeNameMapping[a]; found {
		return name
	} else {
		return "Unknown"
	}
}

var AlertTypeNameMapping = map[AlertType]string{
	GeneralAlert:  "General",
	PreAlarmAlert: "PreAlarm",
	ArmAlert:      "Arm",
	DisarmAlert:   "Disarm",
	FaultAlert:    "Fault",
}

const (
	GeneralAlert  AlertType = 0x00
	PreAlarmAlert AlertType = 0x01
	ArmAlert      AlertType = 0x02
	DisarmAlert   AlertType = 0x03
	FaultAlert    AlertType = 0x04
)

// AlarmType represents the type of alarm, that may result in different sounds or visual patters/colours.
type AlarmType uint16

func (a AlarmType) String() string {
	if name, found := AlarmTypeNameMapping[a]; found {
		return name
	} else {
		return "Unknown"
	}
}

const (
	GeneralAlarm       AlarmType = 0x0000
	FireAlarm          AlarmType = 0x0001
	SecurityAlarm      AlarmType = 0x0002
	EnvironmentalAlarm AlarmType = 0x0003
	PanicAlarm         AlarmType = 0x0004
	TamperAlarm        AlarmType = 0x0005
	DeviceFaultAlarm   AlarmType = 0x0006
	HealthAlarm        AlarmType = 0x0007
)

var AlarmTypeNameMapping = map[AlarmType]string{
	GeneralAlarm:       "General",
	FireAlarm:          "Fire",
	SecurityAlarm:      "Security",
	EnvironmentalAlarm: "Environmental",
	PanicAlarm:         "Panic",
	TamperAlarm:        "Tamper",
	DeviceFaultAlarm:   "DeviceFault",
	HealthAlarm:        "Health",
}

// AlarmWarningDevice is a capability that represents a device that is a warning device or alert mechanism in an alarm system.
type AlarmWarningDevice interface {
	// Alarm turns on warning device.
	Alarm(context.Context, da.Device, AlarmType, float64, bool, time.Duration) error
	// Clear turns off warning device.
	Clear(context.Context, da.Device) error
	// Alert causes a brief warning to be made.
	Alert(context.Context, da.Device, AlarmType, AlertType, float64, bool) error
	// Status returns the state of the warning device.
	Status(context.Context, da.Device) (WarningDeviceState, error)
}

// AlarmWarningDeviceUpdate is sent to inform consumers that an AlarmWarningDevice state has changed.
type AlarmWarningDeviceUpdate struct {
	// Device that has updated state.
	Device da.Device
	// State contains the state of the warning device.
	State WarningDeviceState
}

// WarningDeviceState represents the state of a WarningDevice
type WarningDeviceState struct {
	// Warning represents if the device is currently warning.
	Warning bool
	// AlarmType that is currently being warned about.
	AlarmType AlarmType
	// Volume is the volume of the device, 0..1.0.
	Volume float64
	// Visual is true if a visual alert component is present.
	Visual bool
	// DurationRemaining is the amount of time remaining of this warning.
	DurationRemaining time.Duration
}
