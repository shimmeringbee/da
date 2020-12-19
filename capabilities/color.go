package capabilities

import (
	"context"
	. "github.com/shimmeringbee/da"
	"image/color"
	"time"
)

// ConvertibleColor is an interface that allows acceptance of a color implementation that supports
// modes required by
type ConvertibleColor interface {
	HueSat() (float64, float64)
	XY() (float64, float64)
	Color() color.Color
}

type Color interface {
	// ChangeToTemperature changes the output of the device to be that color over the duration specified.
	ChangeToColor(context.Context, Device, ConvertibleColor, time.Duration) error
	// ChangeToTemperature changes the output of the device to be that temperature over the duration specified.
	ChangeToTemperature(context.Context, Device, float64, time.Duration) error
	// SupportsColor if the device supports outputting a color.
	SupportsColor() bool
	// SupportsTemperature if the device supports outputting a specific white temperature.
	SupportsTemperature() bool
	// Status returns the current status of Color.
	Status(context.Context, Device) (ColorStatus, error)
}

// ColorStatus represents the current color.
type ColorStatus struct {
	// CurrentColor is the current color of the device.
	CurrentColor ConvertibleColor
	// TargetColor is the target color of the device.
	TargetColor ConvertibleColor
	// DurationRemaining is how long until CurrentColor transitions to TargetColor.
	DurationRemaining time.Duration
}
