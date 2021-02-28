package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/stretchr/testify/mock"
)

var _ da.Gateway = (*Gateway)(nil)

type Gateway struct {
	mock.Mock
}

func (m *Gateway) ReadEvent(ctx context.Context) (interface{}, error) {
	args := m.Called(ctx)
	return args.Get(0), args.Error(1)
}

func (m *Gateway) Capability(c da.Capability) interface{} {
	args := m.Called(c)
	return args.Get(0)
}

func (m *Gateway) Capabilities() []da.Capability {
	args := m.Called()
	return args.Get(0).([]da.Capability)
}

func (m *Gateway) Self() da.Device {
	args := m.Called()
	return args.Get(0).(da.Device)
}

func (m *Gateway) Devices() []da.Device {
	args := m.Called()
	return args.Get(0).([]da.Device)
}

func (m *Gateway) Start() error {
	args := m.Called()
	return args.Error(0)
}

func (m *Gateway) Stop() error {
	args := m.Called()
	return args.Error(0)
}
