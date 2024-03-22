package capabilities

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities/color"
	"time"
)

// Light is a capability that provides the ability for a device to control the production of light, often
// a light bulb or strip.
type Light interface {
	// SetBrightness changes the devices current output level to the new required value, over a set duration.
	SetBrightness(context.Context, float64, time.Duration) error
	// SetColor changes the output of the device to be that color over the duration specified.
	SetColor(context.Context, color.ConvertibleColor, time.Duration) error
	// SetTemperature changes the output of the device to be that temperature over the duration specified.
	SetTemperature(context.Context, float64, time.Duration) error

	// SupportsBrightness if the device supports setting of brightness.
	SupportsBrightness(context.Context) (bool, error)
	// SupportsColor if the device supports outputting a color.
	SupportsColor(context.Context) (bool, error)
	// SupportsTemperature if the device supports outputting a specific white temperature.
	SupportsTemperature(context.Context) (bool, error)

	// Status returns the current status of Light.
	Status(context.Context) (LightStatus, error)
}

// Mode represents what the device is set to output.
type Mode uint8

const (
	// ColorMode the device is outputting a color.
	ColorMode Mode = 1
	// TemperatureMode the device is outputting a black body temperature color.
	TemperatureMode Mode = 2
)

// LightStatus represents the current light state.
type LightStatus struct {
	// Current light state.
	Current LightState
	// Target light state.
	Target LightState
	// DurationRemaining is how long until Current transitions to Target.
	DurationRemaining time.Duration
}

type LightState struct {
	// Brightness of light.
	Brightness float64
	// Color of the light.
	Color color.ConvertibleColor
	// Temperature (black body) of light.
	Temperature float64
	// Mode the light is running.
	Mode Mode
}

// LightStatusUpdate is sent to inform consumers of the devices color state, may be sent even if there is no change.
type LightStatusUpdate struct {
	// Device that is informing of it's on off state.
	Device da.Device
	// New state of device.
	State LightStatus
}
