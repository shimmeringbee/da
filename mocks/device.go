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

type SimpleDevice struct {
	SGateway      da.Gateway
	SIdentifier   da.Identifier
	SCapabilities []da.Capability
}

func (s SimpleDevice) Gateway() da.Gateway {
	return s.SGateway
}

func (s SimpleDevice) Identifier() da.Identifier {
	return s.SIdentifier
}

func (s SimpleDevice) Capabilities() []da.Capability {
	return s.SCapabilities
}

func (s SimpleDevice) Capability(_ da.Capability) da.BasicCapability {
	panic("use MockDevice instead")
}

var _ da.Device = (*SimpleDevice)(nil)
