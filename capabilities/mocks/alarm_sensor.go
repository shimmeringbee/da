package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
)

var _ capabilities.AlarmSensor = (*AlarmSensor)(nil)

type AlarmSensor struct {
	mock.Mock
}

func (m *AlarmSensor) Status(c context.Context, d da.Device) (map[capabilities.SensorType]bool, error) {
	args := m.Called(c, d)
	return args.Get(0).(map[capabilities.SensorType]bool), args.Error(1)
}
