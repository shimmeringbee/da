package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
	"time"
)

var _ capabilities.AlarmWarningDevice = (*AlarmWarningDevice)(nil)

type AlarmWarningDevice struct {
	mock.Mock
}

func (m *AlarmWarningDevice) Alarm(c context.Context, d da.Device, a capabilities.AlarmType, vol float64, vis bool, dur time.Duration) error {
	args := m.Called(c, d, a, vol, vis, dur)
	return args.Error(0)
}

func (m *AlarmWarningDevice) Clear(c context.Context, d da.Device) error {
	args := m.Called(c, d)
	return args.Error(0)
}

func (m *AlarmWarningDevice) Alert(c context.Context, d da.Device, alarm capabilities.AlarmType, alert capabilities.AlertType, vol float64, vis bool) error {
	args := m.Called(c, d, alarm, alert, vol, vis)
	return args.Error(0)
}

func (m *AlarmWarningDevice) Status(c context.Context, d da.Device) (capabilities.WarningDeviceState, error) {
	args := m.Called(c, d)
	return args.Get(0).(capabilities.WarningDeviceState), args.Error(1)
}
