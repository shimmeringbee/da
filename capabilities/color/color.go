package color

import (
	"image/color"
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
