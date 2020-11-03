package capabilities

import (
	. "github.com/shimmeringbee/da"
)

const (
	/* Basic capabilities to permit management of devices. */
	DeviceDiscoveryFlag       = Capability(0x0000)
	EnumerateDeviceFlag       = Capability(0x0001)
	HasProductInformationFlag = Capability(0x0002)
	DeviceRemovalFlag         = Capability(0x0003)
	PowerStatusFlag           = Capability(0x0100)
	ConnectivityFlag          = Capability(0x0101)
	DeviceAlarmsFlag          = Capability(0x0102)

	/* Capabilities to change the physical state of devices. */
	OnOffFlag            = Capability(0x1000)
	LevelFlag            = Capability(0x1001)
	ColorTemperatureFlag = Capability(0x1002)

	/* Capabilities which humans interact with to provide input. */
	BasicControllerFlag = Capability(0x1200)

	/* Capabilities to read information from the environment. */
	TemperatureSensorFlag      = Capability(0x2000)
	RelativeHumiditySensorFlag = Capability(0x2001)
	PressureSensorFlag         = Capability(0x2002)
	PresenceSensorFlag         = Capability(0x2003)
	IlluminationSensorFlag     = Capability(0x2004)

	/* Capabilities to read information from power systems. */
	EnergyMeasurementFlag = Capability(0x2100)

	/* Alarm Systems. */
	AlarmSensorFlag        = Capability(0x3000)
	AlarmWarningDeviceFlag = Capability(0x3001)

	/* Developer Functionality. */
	LocalDebugFlag          = Capability(0xff00)
	RemoteDebugFlag         = Capability(0xff01)
	MessageCaptureDebugFlag = Capability(0xff20)
)

var StandardNames = map[Capability]string{
	DeviceDiscoveryFlag:        "DeviceDiscovery",
	EnumerateDeviceFlag:        "EnumerateDevice",
	HasProductInformationFlag:  "HasProductInformation",
	DeviceRemovalFlag:          "DeviceRemoval",
	PowerStatusFlag:            "PowerStatus",
	ConnectivityFlag:           "Connectivity",
	DeviceAlarmsFlag:           "DeviceAlarms",
	OnOffFlag:                  "OnOff",
	LevelFlag:                  "Level",
	ColorTemperatureFlag:       "ColorTemperature",
	BasicControllerFlag:        "BasicController",
	TemperatureSensorFlag:      "TemperatureSensor",
	RelativeHumiditySensorFlag: "RelativeHumiditySensor",
	PressureSensorFlag:         "PressureSensor",
	PresenceSensorFlag:         "PresenceSensor",
	IlluminationSensorFlag:     "IlluminationSensor",
	EnergyMeasurementFlag:      "EnergyMeasurement",
	AlarmSensorFlag:            "AlarmSensor",
	AlarmWarningDeviceFlag:     "AlarmWarningDevice",
	LocalDebugFlag:             "LocalDebug",
	RemoteDebugFlag:            "RemoteDebug",
	MessageCaptureDebugFlag:    "MessageCaptureDebug",
}
