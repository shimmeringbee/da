package mocks

import (
	"github.com/shimmeringbee/da"
	"github.com/stretchr/testify/mock"
)

type MockDevice struct {
	mock.Mock
}

func (m *MockDevice) Gateway() da.Gateway {
	return m.Called().Get(0).(da.Gateway)
}

func (m *MockDevice) Identifier() da.Identifier {
	return m.Called().Get(0).(da.Identifier)
}

func (m *MockDevice) Capabilities() []da.Capability {
	return m.Called().Get(0).([]da.Capability)
}

func (m *MockDevice) Capability(capability da.Capability) da.BasicCapability {
	return m.Called(capability).Get(0).(da.BasicCapability)
}

var _ da.Device = (*MockDevice)(nil)
