package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
)

var _ capabilities.RelativeHumiditySensor = (*RelativeHumiditySensor)(nil)

type RelativeHumiditySensor struct {
	mock.Mock
}

func (m *RelativeHumiditySensor) Reading(c context.Context) ([]capabilities.RelativeHumidityReading, error) {
	args := m.Called(c)
	return args.Get(0).([]capabilities.RelativeHumidityReading), args.Error(1)
}

func (m *RelativeHumiditySensor) Capability() da.Capability {
	args := m.Called()
	return args.Get(0).(da.Capability)
}

func (m *RelativeHumiditySensor) Name() string {
	args := m.Called()
	return args.String(0)
}
