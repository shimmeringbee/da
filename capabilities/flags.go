package capabilities

import (
	"github.com/shimmeringbee/da"
)

const (
	/* Basic capabilities to permit management of devices. */
	DeviceDiscoveryFlag       = da.Capability(0x0000)
	EnumerateDeviceFlag       = da.Capability(0x0001)
	HasProductInformationFlag = da.Capability(0x0002)
	DeviceRemovalFlag         = da.Capability(0x0003)
	PowerSupplyFlag           = da.Capability(0x0100)
	ConnectivityFlag          = da.Capability(0x0101)
	DeviceAlarmsFlag          = da.Capability(0x0102)

	/* Capabilities to change the physical state of devices. */
	OnOffFlag = da.Capability(0x1000)
	LevelFlag = da.Capability(0x1001)
	ColorFlag = da.Capability(0x1002)

	/* Capabilities which humans interact with to provide input. */
	BasicControllerFlag = da.Capability(0x1200)

	/* Capabilities to read information from the environment. */
	TemperatureSensorFlag      = da.Capability(0x2000)
	RelativeHumiditySensorFlag = da.Capability(0x2001)
	PressureSensorFlag         = da.Capability(0x2002)
	OccupancySensorFlag        = da.Capability(0x2003)
	IlluminationSensorFlag     = da.Capability(0x2004)

	/* Capabilities to read information from power systems. */
	EnergyMeasurementFlag = da.Capability(0x2100)

	/* Alarm Systems. */
	AlarmSensorFlag        = da.Capability(0x3000)
	AlarmWarningDeviceFlag = da.Capability(0x3001)

	/* Developer Functionality. */
	LocalDebugFlag          = da.Capability(0xff00)
	RemoteDebugFlag         = da.Capability(0xff01)
	MessageCaptureDebugFlag = da.Capability(0xff20)
)

var StandardNames = map[da.Capability]string{
	DeviceDiscoveryFlag:        "DeviceDiscovery",
	EnumerateDeviceFlag:        "EnumerateDevice",
	HasProductInformationFlag:  "HasProductInformation",
	DeviceRemovalFlag:          "DeviceRemoval",
	PowerSupplyFlag:            "PowerSupply",
	ConnectivityFlag:           "Connectivity",
	DeviceAlarmsFlag:           "DeviceAlarms",
	OnOffFlag:                  "OnOff",
	LevelFlag:                  "Level",
	ColorFlag:                  "Color",
	BasicControllerFlag:        "BasicController",
	TemperatureSensorFlag:      "TemperatureSensor",
	RelativeHumiditySensorFlag: "RelativeHumiditySensor",
	PressureSensorFlag:         "PressureSensor",
	OccupancySensorFlag:        "OccupancySensor",
	IlluminationSensorFlag:     "IlluminationSensor",
	EnergyMeasurementFlag:      "EnergyMeasurement",
	AlarmSensorFlag:            "AlarmSensor",
	AlarmWarningDeviceFlag:     "AlarmWarningDevice",
	LocalDebugFlag:             "LocalDebug",
	RemoteDebugFlag:            "RemoteDebug",
	MessageCaptureDebugFlag:    "MessageCaptureDebug",
}
