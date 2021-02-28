package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
	"time"
)

var _ capabilities.Level = (*Level)(nil)

type Level struct {
	mock.Mock
}

func (m *Level) Status(c context.Context, d da.Device) (capabilities.LevelStatus, error) {
	args := m.Called(c, d)
	return args.Get(0).(capabilities.LevelStatus), args.Error(1)
}

func (m *Level) Change(c context.Context, d da.Device, l float64, t time.Duration) error {
	args := m.Called(c, d, l, t)
	return args.Error(0)
}
