package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
)

var _ capabilities.TemperatureSensor = (*TemperatureSensor)(nil)

type TemperatureSensor struct {
	mock.Mock
}

func (m *TemperatureSensor) Reading(c context.Context, d da.Device) ([]capabilities.TemperatureReading, error) {
	args := m.Called(c, d)
	return args.Get(0).([]capabilities.TemperatureReading), args.Error(1)
}
