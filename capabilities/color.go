package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities/color"
	"time"
)

// Color is a capability that provides the ability for a device to set the color of an output, this
// is usually a bulb or light strip.
type Color interface {
	// ChangeTemperature changes the output of the device to be that color over the duration specified.
	ChangeColor(context.Context, da.Device, color.ConvertibleColor, time.Duration) error
	// ChangeTemperature changes the output of the device to be that temperature over the duration specified.
	ChangeTemperature(context.Context, da.Device, float64, time.Duration) error
	// SupportsColor if the device supports outputting a color.
	SupportsColor() bool
	// SupportsTemperature if the device supports outputting a specific white temperature.
	SupportsTemperature() bool
	// Status returns the current status of Color.
	Status(context.Context, da.Device) (ColorStatus, error)
}

// Mode represents what the device is set to output.
type Mode uint8

const (
	// ColorMode the device is outputting a color.
	ColorMode Mode = 1
	// TemperatureMode the device is outputting a black body temperature color.
	TemperatureMode Mode = 2
)

// ColorStatus represents the current color.
type ColorStatus struct {
	// Mode is the current color mode.
	Mode Mode
	// Color information if the Mode is ColorMode.
	Color ColorSettings
	// Temperature information if the Mode is TemperatureMode.
	Temperature TemperatureSettings
	// DurationRemaining is how long until Current transitions to Target.
	DurationRemaining time.Duration
}

// ColorSettings holds the current color and target color.
type ColorSettings struct {
	// Current color of the device.
	Current color.ConvertibleColor
	// Target color of the device.
	Target color.ConvertibleColor
}

// TemperatureSettings holds the current color temperature and target color temperature.
type TemperatureSettings struct {
	// Current temperature of the device.
	Current float64
	// Target temperature of the device.
	Target float64
}

// ColorStatusUpdate is sent to inform consumers of the devices color state, may be sent even if there is no change.
type ColorStatusUpdate struct {
	// Device that is informing of it's on off state.
	Device da.Device
	// New state of device.
	State ColorStatus
}
