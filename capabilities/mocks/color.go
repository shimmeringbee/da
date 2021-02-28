package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/shimmeringbee/da/capabilities/color"
	"github.com/stretchr/testify/mock"
	"time"
)

var _ capabilities.Color = (*Color)(nil)

type Color struct {
	mock.Mock
}

func (m *Color) ChangeColor(ctx context.Context, d da.Device, color color.ConvertibleColor, duration time.Duration) error {
	args := m.Called(ctx, d, color, duration)
	return args.Error(0)
}

func (m *Color) ChangeTemperature(ctx context.Context, d da.Device, f float64, duration time.Duration) error {
	args := m.Called(ctx, d, f, duration)
	return args.Error(0)
}

func (m *Color) SupportsColor(ctx context.Context, device da.Device) (bool, error) {
	args := m.Called(ctx, device)
	return args.Bool(0), args.Error(1)
}

func (m *Color) SupportsTemperature(ctx context.Context, device da.Device) (bool, error) {
	args := m.Called(ctx, device)
	return args.Bool(0), args.Error(1)
}

func (m *Color) Status(ctx context.Context, d da.Device) (capabilities.ColorStatus, error) {
	args := m.Called(ctx, d)
	return args.Get(0).(capabilities.ColorStatus), args.Error(1)
}
