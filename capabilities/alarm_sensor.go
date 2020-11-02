package capabilities

import (
	"context"
	"fmt"
	"github.com/shimmeringbee/da"
)

// SensorType provides information about the type of sensors on an Alarm Sensor.
type SensorType uint16

// String returns a string representation of the alarm.
func (z SensorType) String() string {
	if name, found := NameMapping[z]; found {
		return name
	} else {
		return fmt.Sprintf("Unknown")
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
	DeviceTest           SensorType = 0xf010
)

var NameMapping = map[SensorType]string{
	General:                "General",
	GeneralEmergency:       "GeneralEmergency",
	GeneralWarningDevice:   "GeneralWarningDevice",
	FireTemperature:        "FireTemperature",
	FireIonizing:           "FireIonizing",
	FirePhotoelectric:      "FirePhotoelectric",
	FireBreakGlass:         "FireBreakGlass",
	FirePreAlarm:           "FirePreAlarm",
	FireOther:              "FireOther",
	GasCarbonMonoxide:      "GasCarbonMonoxide",
	GasCarbonDioxide:       "GasCarbonDioxide",
	GasOxygen:              "GasOxygen",
	GasCombustible:         "GasCombustible",
	GasRadon:               "GasRadon",
	Radiation:              "Radiation",
	SecurityContact:        "SecurityContact",
	SecurityMotion:         "SecurityMotion",
	SecurityVibration:      "SecurityVibration",
	SecurityGlassBreak:     "SecurityGlassBreak",
	SecurityOther:          "SecurityOther",
	SecurityInfrastructure: "SecurityInfrastructure",
	SecurityPanic:          "SecurityPanic",
	SecurityKeypad:         "SecurityKeypad",
	HealthFall:             "HealthFall",
	DeviceTamper:           "DeviceTamper",
	DeviceBatteryLow:       "DeviceBatteryLow",
	DeviceBatteryFailure:   "DeviceBatteryFailure",
	DeviceMainsFailure:     "DeviceMainsFailure",
	DeviceFailure:          "DeviceFailure",
	DeviceTest:             "DeviceTest",
}

// AlarmSensor is a capability that represents a device that is a sensor in an alarm system.
type AlarmSensor interface {
	// Status returns all states on this sensor.
	Status(context.Context, da.Device) (map[SensorType]bool, error)
}

// AlarmSensorUpdate is sent to inform consumers that an AlarmSensors state has changed.
type AlarmSensorUpdate struct {
	// Device that has updated state.
	Device da.Device
	// States contains a map of all alarm states, including those that have not changed.
	States map[SensorType]bool
}
