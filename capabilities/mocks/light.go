package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/shimmeringbee/da/capabilities/color"
	"github.com/stretchr/testify/mock"
	"time"
)

var _ capabilities.Light = (*Light)(nil)

type Light struct {
	mock.Mock
}

func (l *Light) SetBrightness(c context.Context, f float64, duration time.Duration) error {
	args := l.Called(c, f, duration)
	return args.Error(0)
}

func (l *Light) SetColor(c context.Context, color color.ConvertibleColor, duration time.Duration) error {
	args := l.Called(c, color, duration)
	return args.Error(0)
}

func (l *Light) SetTemperature(c context.Context, f float64, duration time.Duration) error {
	args := l.Called(c, f, duration)
	return args.Error(0)
}

func (l *Light) SupportsBrightness(c context.Context) (bool, error) {
	args := l.Called(c)
	return args.Bool(0), args.Error(1)
}

func (l *Light) SupportsColor(c context.Context) (bool, error) {
	args := l.Called(c)
	return args.Bool(0), args.Error(1)
}

func (l *Light) SupportsTemperature(c context.Context) (bool, error) {
	args := l.Called(c)
	return args.Bool(0), args.Error(1)
}

func (l *Light) Status(c context.Context) (capabilities.LightStatus, error) {
	args := l.Called(c)
	return args.Get(0).(capabilities.LightStatus), args.Error(1)
}

func (l *Light) Capability() da.Capability {
	args := l.Called()
	return args.Get(0).(da.Capability)
}

func (l *Light) Name() string {
	args := l.Called()
	return args.String(0)
}
