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

func (m *AlarmWarningDevice) Alarm(c context.Context, a capabilities.AlarmType, vol float64, vis bool, dur time.Duration) error {
	args := m.Called(c, a, vol, vis, dur)
	return args.Error(0)
}

func (m *AlarmWarningDevice) Clear(c context.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *AlarmWarningDevice) Alert(c context.Context, alarm capabilities.AlarmType, alert capabilities.AlertType, vol float64, vis bool) error {
	args := m.Called(c, alarm, alert, vol, vis)
	return args.Error(0)
}

func (m *AlarmWarningDevice) Status(c context.Context) (capabilities.WarningDeviceState, error) {
	args := m.Called(c)
	return args.Get(0).(capabilities.WarningDeviceState), args.Error(1)
}

func (m *AlarmWarningDevice) Capability() da.Capability {
	args := m.Called()
	return args.Get(0).(da.Capability)
}

func (m *AlarmWarningDevice) Name() string {
	args := m.Called()
	return args.String(0)
}
