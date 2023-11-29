package color

import "github.com/lucasb-eyer/go-colorful"

// NativeColorspace describes the native color space of a convertible color.
type NativeColorspace uint8

const (
	// HueSat the color is represented as Hue and saturation of the sRGB colorspace.
	HueSat NativeColorspace = 0
	// XYY the color is represented as the XYY coordinates fo the CIE 1931 colorspace.
	XYY NativeColorspace = 1
	// SRGB the color is represented as red, green and blue in the screen colorspace.
	SRGB NativeColorspace = 2
)

// ConvertibleColor is an interface that allows acceptance of a color implementation that supports
// modes required by a device.
type ConvertibleColor interface {
	// HSV returns Hue and Saturation, converted from an internal format if required.
	HSV() (float64, float64, float64)
	// XYY returns X and Y in the CIE 1931 colour space, converted from an internal format if required.
	XYY() (float64, float64, float64)
	// RGB returns colour using 8 bit values, converted from an internal format if required.
	RGB() (uint8, uint8, uint8)
	// NativeColorspace returns the native colour space of the colour.
	NativeColorspace() NativeColorspace
}

var _ ConvertibleColor = (*XYColor)(nil)

type XYColor struct {
	X  float64
	Y  float64
	Y2 float64
}

func (c XYColor) HSV() (float64, float64, float64) {
	return colorful.Xyy(c.X, c.Y, c.Y2).Hsv()
}

func (c XYColor) XYY() (float64, float64, float64) {
	return c.X, c.Y, c.Y2
}

func (c XYColor) RGB() (uint8, uint8, uint8) {
	return colorful.Xyy(c.X, c.Y, 100.0).RGB255()
}

func (c XYColor) NativeColorspace() NativeColorspace {
	return XYY
}

var _ ConvertibleColor = (*HSVColor)(nil)

type HSVColor struct {
	Hue   float64
	Sat   float64
	Value float64
}

func (c HSVColor) HSV() (float64, float64, float64) {
	return c.Hue, c.Sat, c.Value
}

func (c HSVColor) XYY() (float64, float64, float64) {
	return colorful.Hsv(c.Hue, c.Sat, c.Value).Xyy()
}

func (c HSVColor) RGB() (uint8, uint8, uint8) {
	return colorful.Hsv(c.Hue, c.Sat, 1.0).RGB255()
}

func (c HSVColor) NativeColorspace() NativeColorspace {
	return HueSat
}

var _ ConvertibleColor = (*SRGBColor)(nil)

type SRGBColor struct {
	R uint8
	G uint8
	B uint8
}

func (c SRGBColor) HSV() (float64, float64, float64) {
	return colorful.Color{R: float64(c.R) / 255.0, G: float64(c.G) / 255.0, B: float64(c.B) / 255.0}.Hsv()
}

func (c SRGBColor) XYY() (float64, float64, float64) {
	return colorful.Color{R: float64(c.R) / 255.0, G: float64(c.G) / 255.0, B: float64(c.B) / 255.0}.Xyy()
}

func (c SRGBColor) RGB() (uint8, uint8, uint8) {
	return c.R, c.G, c.B
}

func (c SRGBColor) NativeColorspace() NativeColorspace {
	return SRGB
}
