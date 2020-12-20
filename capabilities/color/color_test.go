package color

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_XYColor(t *testing.T) {
	t.Run("NativeColorspace returns XYY", func(t *testing.T) {
		color := XYColor{X: 0.5, Y: 0.75}

		expectedNativeColorSpace := XYY
		actualNativeColorSpace := color.NativeColorspace()

		assert.Equal(t, expectedNativeColorSpace, actualNativeColorSpace)
	})

	t.Run("can return XYY CIE 1931", func(t *testing.T) {
		color := XYColor{X: 0.5, Y: 0.75, Y2: 0.5}

		rX, rY, rY2 := color.XYY()

		assert.Equal(t, color.X, rX)
		assert.Equal(t, color.Y, rY)
		assert.Equal(t, color.Y2, rY2)
	})

	t.Run("can return Hue Sat", func(t *testing.T) {
		color := XYColor{X: 0.4183784094737513, Y: 0.3946924771203395, Y2: 0.6018037581822194}

		rHue, rSat, rVal := color.HSV()

		assert.InDelta(t, 30, rHue, 0.01)
		assert.InDelta(t, 0.499, rSat, 0.01)
		assert.InDelta(t, 1.0, rVal, 0.01)
	})

	t.Run("can return Color", func(t *testing.T) {
		color := XYColor{X: 0.4183784094737513, Y: 0.3946924771203395, Y2: 0.6018037581822194}

		r, g, b := color.RGB()

		assert.Equal(t, 203, int(r))
		assert.Equal(t, 178, int(g))
		assert.Equal(t, 153, int(b))
	})
}

func Test_HueSatColor(t *testing.T) {
	t.Run("NativeColorspace returns HSV", func(t *testing.T) {
		color := HSVColor{Hue: 0.5, Sat: 0.75, Value: 1.0}

		expectedNativeColorSpace := HueSat
		actualNativeColorSpace := color.NativeColorspace()

		assert.Equal(t, expectedNativeColorSpace, actualNativeColorSpace)
	})

	t.Run("can return XYY CIE 1931", func(t *testing.T) {
		color := HSVColor{Hue: 40.21, Sat: 0.749, Value: 1.0}

		rX, rY, rY2 := color.XYY()

		assert.InDelta(t, 0.4573, rX, 0.0001)
		assert.InDelta(t, 0.4446, rY, 0.0001)
		assert.InDelta(t, 0.5934, rY2, 0.0001)
	})

	t.Run("can return Hue Sat", func(t *testing.T) {
		color := HSVColor{Hue: 0.5, Sat: 0.75, Value: 0.45}

		rHue, rSat, rVal := color.HSV()

		assert.Equal(t, color.Hue, rHue)
		assert.Equal(t, color.Sat, rSat)
		assert.Equal(t, color.Value, rVal)
	})

	t.Run("can return Color", func(t *testing.T) {
		color := HSVColor{Hue: 40.21, Sat: 0.749, Value: 1.0}

		r, g, b := color.RGB()

		assert.Equal(t, uint8(255), r)
		assert.Equal(t, uint8(192), g)
		assert.Equal(t, uint8(64), b)

		rgbColor := SRGBColor{R: r, G: g, B: b}
		rHue, rSat, rVal := rgbColor.HSV()

		assert.InDelta(t, color.Hue, rHue, 0.001)
		assert.InDelta(t, color.Sat, rSat, 0.001)
		assert.InDelta(t, color.Value, rVal, 0.001)
	})
}

func Test_SRGBColor(t *testing.T) {
	t.Run("NativeColorspace returns XYY", func(t *testing.T) {
		color := SRGBColor{R: 192, G: 128, B: 64}

		expectedNativeColorSpace := SRGB
		actualNativeColorSpace := color.NativeColorspace()

		assert.Equal(t, expectedNativeColorSpace, actualNativeColorSpace)
	})

	t.Run("can return XYY CIE 1931", func(t *testing.T) {
		color := SRGBColor{R: 192, G: 128, B: 64}

		rX, rY, rY2 := color.XYY()

		assert.InDelta(t, 0.4613, rX, 0.0001)
		assert.InDelta(t, 0.4101, rY, 0.0001)
		assert.InDelta(t, 0.2702, rY2, 0.0001)
	})

	t.Run("can return Hue Sat for red", func(t *testing.T) {
		color := SRGBColor{R: 255, G: 0, B: 0}

		rHue, rSat, rVal := color.HSV()

		assert.Equal(t, 0.0, rHue)
		assert.Equal(t, 1.0, rSat)
		assert.Equal(t, 1.0, rVal)
	})

	t.Run("can return Hue Sat for blue", func(t *testing.T) {
		color := SRGBColor{R: 0, G: 0, B: 255}

		rHue, rSat, rVal := color.HSV()

		assert.Equal(t, 240.0, rHue)
		assert.Equal(t, 1.0, rSat)
		assert.Equal(t, 1.0, rVal)
	})

	t.Run("can return Hue Sat for white", func(t *testing.T) {
		color := SRGBColor{R: 255, G: 255, B: 255}

		rHue, rSat, rVal := color.HSV()

		assert.Equal(t, 0.0, rHue)
		assert.Equal(t, 0.0, rSat)
		assert.Equal(t, 1.0, rVal)
	})

	t.Run("can return Hue Sat for cyan", func(t *testing.T) {
		color := SRGBColor{R: 0, G: 255, B: 255}

		rHue, rSat, rVal := color.HSV()

		assert.Equal(t, 180.0, rHue)
		assert.Equal(t, 1.0, rSat)
		assert.Equal(t, 1.0, rVal)
	})

	t.Run("can return Hue Sat for lime green", func(t *testing.T) {
		color := SRGBColor{R: 122, G: 255, B: 255}

		rHue, rSat, rVal := color.HSV()

		assert.Equal(t, 180.0, rHue)
		assert.InDelta(t, 0.521, rSat, 0.001)
		assert.Equal(t, 1.0, rVal)
	})

	t.Run("can return RGB", func(t *testing.T) {
		color := SRGBColor{R: 192, G: 128, B: 64}

		r, g, b := color.RGB()

		assert.Equal(t, color.R, r)
		assert.Equal(t, color.G, g)
		assert.Equal(t, color.B, b)
	})
}
