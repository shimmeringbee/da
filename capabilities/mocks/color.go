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

func (m *Color) ChangeColor(ctx context.Context, color color.ConvertibleColor, duration time.Duration) error {
	args := m.Called(ctx, color, duration)
	return args.Error(0)
}

func (m *Color) ChangeTemperature(ctx context.Context, f float64, duration time.Duration) error {
	args := m.Called(ctx, f, duration)
	return args.Error(0)
}

func (m *Color) SupportsColor(ctx context.Context) (bool, error) {
	args := m.Called(ctx)
	return args.Bool(0), args.Error(1)
}

func (m *Color) SupportsTemperature(ctx context.Context) (bool, error) {
	args := m.Called(ctx)
	return args.Bool(0), args.Error(1)
}

func (m *Color) Status(ctx context.Context) (capabilities.ColorStatus, error) {
	args := m.Called(ctx)
	return args.Get(0).(capabilities.ColorStatus), args.Error(1)
}

func (m *Color) Capability() da.Capability {
	args := m.Called()
	return args.Get(0).(da.Capability)
}

func (m *Color) Name() string {
	args := m.Called()
	return args.String(0)
}
