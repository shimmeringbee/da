package capabilities

import (
	. "github.com/shimmeringbee/da"
)

const (
	/* Basic capabilities to permit management of devices. */
	DeviceDiscoveryFlag       = Capability(0x0000)
	EnumerateDeviceFlag       = Capability(0x0001)
	HasProductInformationFlag = Capability(0x0002)
	PowerStatusFlag           = Capability(0x0100)
	ConnectivityFlag          = Capability(0x0101)
	DeviceAlarmsFlag          = Capability(0x0102)

	/* Capabilities to change the physical state of devices. */
	OnOffFlag            = Capability(0x1000)
	LevelFlag            = Capability(0x1001)
	ColorTemperatureFlag = Capability(0x1002)

	/* Capabilities to read information from the environment. */
	TemperatureSensorFlag  = Capability(0x2000)
	HumiditySensorFlag     = Capability(0x2001)
	PressureSensorFlag     = Capability(0x2002)
	PresenceSensorFlag     = Capability(0x2003)
	IlluminationSensorFlag = Capability(0x2004)

	/* Capabilities to read information from power systems. */
	EnergyMeasurementFlag = Capability(0x2100)

	/* Alarm Systems. */
	AlarmZoneFlag          = Capability(0x3000)
	AlarmWarningDeviceFlag = Capability(0x3001)

	/* Developer Functionality. */
	LocalDebugFlag          = Capability(0xff00)
	RemoteDebugFlag         = Capability(0xff01)
	MessageCaptureDebugFlag = Capability(0xff20)
)
