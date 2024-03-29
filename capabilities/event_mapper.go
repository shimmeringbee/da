package capabilities

import "github.com/shimmeringbee/da"

// EventToCapability maps a device abstraction capabilities event message back to the capability flag.
func EventToCapability(v interface{}) (da.Capability, bool) {
	switch v.(type) {
	case AlarmSensorUpdate:
		return AlarmSensorFlag, true
	case AlarmWarningDeviceUpdate:
		return AlarmWarningDeviceFlag, true
	case LightStatusUpdate:
		return LightFlag, true
	case DeviceDiscoveryEnabled, DeviceDiscoveryDisabled:
		return DeviceDiscoveryFlag, true
	case EnumerateDeviceStart, EnumerateDeviceStopped:
		return EnumerateDeviceFlag, true
	case IlluminationSensorState:
		return IlluminationSensorFlag, true
	case LocalDebugStart, LocalDebugSuccess, LocalDebugFailure:
		return LocalDebugFlag, true
	case MessageCaptureStart, MessageCapture, MessageCaptureStop:
		return MessageCaptureDebugFlag, true
	case OccupancySensorState:
		return OccupancySensorFlag, true
	case OnOffState:
		return OnOffFlag, true
	case PowerStatusUpdate:
		return PowerSupplyFlag, true
	case PressureSensorState:
		return PressureSensorFlag, true
	case RelativeHumiditySensorState:
		return RelativeHumiditySensorFlag, true
	case RemoteDebugStart, RemoteDebugSuccess, RemoteDebugFailure:
		return RemoteDebugFlag, true
	case TemperatureSensorState:
		return TemperatureSensorFlag, true
	default:
		return 0, false
	}
}
