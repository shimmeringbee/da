package capabilities

import (
	"context"
	. "github.com/shimmeringbee/da"
	"image/color"
	"time"
)

// NativeColorspace describes the native color space of a convertible color.
type NativeColorspace uint8

const (
	// HueSat the color is represented as hue and saturation of the sRGB colorspace.
	HueSat NativeColorspace = 0
	// XY the color is represented as the XY coordinates fo the CIE 1931 colorspace.
	XY NativeColorspace = 1
	// SRGB the color is represented as red, green and blue in the screen colorspace.
	SRGB NativeColorspace = 2
)

// ConvertibleColor is an interface that allows acceptance of a color implementation that supports
// modes required by a device.
type ConvertibleColor interface {
	// HueSat returns Hue and Saturation, converted from an internal format if required.
	HueSat() (float64, float64)
	// XY returns X and Y in the CIE 1931 colour space, converted from an internal format if required.
	XY() (float64, float64)
	// Color returns a Go color representing sRGB, converted from an internal format if required.
	Color() color.Color
	// NativeColorspace returns the native colour space of the colour.
	NativeColorspace() NativeColorspace
}

// Color is a capability that provides the ability for a device to set the color of an output, this
// is usually a bulb or light strip.
type Color interface {
	// ChangeTemperature changes the output of the device to be that color over the duration specified.
	ChangeColor(context.Context, Device, ConvertibleColor, time.Duration) error
	// ChangeTemperature changes the output of the device to be that temperature over the duration specified.
	ChangeTemperature(context.Context, Device, float64, time.Duration) error
	// SupportsColor if the device supports outputting a color.
	SupportsColor() bool
	// SupportsTemperature if the device supports outputting a specific white temperature.
	SupportsTemperature() bool
	// Status returns the current status of Color.
	Status(context.Context, Device) (ColorStatus, error)
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
	Color struct {
		// Current color of the device.
		Current ConvertibleColor
		// Target color of the device.
		Target ConvertibleColor
	}
	// Temperature information if the Mode is TemperatureMode.
	Temperature struct {
		// Current temperature of the device.
		Current float64
		// Target temperature of the device.
		Target float64
	}
	// DurationRemaining is how long until Current transitions to Target.
	DurationRemaining time.Duration
}
