package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
)

// SensorType provides information about the type of sensors on an Alarm Sensor.
type SensorType uint16

// String returns a string representation of the alarm.
func (z SensorType) String() string {
	switch z {
	case General:
		return "General"
	case GeneralEmergency:
		return "GeneralEmergency"
	case GeneralWarningDevice:
		return "GeneralWarningDevice"
	case FireTemperature:
		return "FireTemperature"
	case FireIonizing:
		return "FireIonizing"
	case FirePhotoelectric:
		return "FirePhotoelectric"
	case FireBreakGlass:
		return "FireBreakGlass"
	case FirePreAlarm:
		return "FirePreAlarm"
	case FireOther:
		return "FireOther"
	case GasCarbonMonoxide:
		return "GasCarbonMonoxide"
	case GasCarbonDioxide:
		return "GasCarbonDioxide"
	case GasOxygen:
		return "GasOxygen"
	case GasCombustible:
		return "GasCombustible"
	case GasRadon:
		return "GasRadon"
	case Radiation:
		return "Radiation"
	case SecurityContact:
		return "SecurityContact"
	case SecurityMotion:
		return "SecurityMotion"
	case SecurityVibration:
		return "SecurityVibration"
	case SecurityGlassBreak:
		return "SecurityGlassBreak"
	case SecurityOther:
		return "SecurityOther"
	case SecurityInfrastructure:
		return "SecurityInfrastructure"
	case SecurityPanic:
		return "SecurityPanic"
	case SecurityKeypad:
		return "SecurityKeypad"
	case HealthFall:
		return "HealthFall"
	case DeviceTamper:
		return "DeviceTamper"
	case DeviceBatteryLow:
		return "DeviceBatteryLow"
	case DeviceBatteryFailure:
		return "DeviceBatteryFailure"
	case DeviceMainsFailure:
		return "DeviceMainsFailure"
	case DeviceFailure:
		return "DeviceFailure"
	default:
		return "Unknown"
	}
}

const (
	General              SensorType = 0x0000
	GeneralEmergency     SensorType = 0x0001
	GeneralWarningDevice SensorType = 0x0010

	FireTemperature   SensorType = 0x1000
	FireIonizing      SensorType = 0x1001
	FirePhotoelectric SensorType = 0x1002
	FireBreakGlass    SensorType = 0x1003
	FirePreAlarm      SensorType = 0x1f00
	FireOther         SensorType = 0x1fff

	GasCarbonMonoxide SensorType = 0x2000
	GasCarbonDioxide  SensorType = 0x2001
	GasOxygen         SensorType = 0x2002
	GasCombustible    SensorType = 0x2003
	GasRadon          SensorType = 0x2004

	Radiation SensorType = 0x2100

	SecurityContact    SensorType = 0x3000
	SecurityMotion     SensorType = 0x3001
	SecurityVibration  SensorType = 0x3002
	SecurityGlassBreak SensorType = 0x3003
	SecurityOther      SensorType = 0x3fff

	SecurityInfrastructure SensorType = 0x3100
	SecurityPanic          SensorType = 0x3101
	SecurityKeypad         SensorType = 0x3102

	HealthFall SensorType = 0x4000

	DeviceTamper         SensorType = 0xf000
	DeviceBatteryLow     SensorType = 0xf001
	DeviceBatteryFailure SensorType = 0xf002
	DeviceMainsFailure   SensorType = 0xf003
	DeviceFailure        SensorType = 0xf004
)

// AlarmSensor is a capability that represents a device that is a sensor in an alarm system.
type AlarmSensor interface {
	// Status returns all states on this sensor.
	Status(context.Context, da.Device) ([]AlarmSensorState, error)
}

// AlarmSensorState represents an individual alarm sensor state on a alarm sensor.
type AlarmSensorState struct {
	// SensorType is the type of sensor that is changing
	SensorType SensorType
	// InAlarm is true if that sensor has triggered.
	InAlarm bool
}

// AlarmSensorUpdate is sent to inform consumers that an AlarmSensors state has changed.
type AlarmSensorUpdate struct {
	// Device that has updated state.
	Device da.Device
	// States contains a list of all alarm states, including those that have not changed.
	States []AlarmSensorState
}
