package mocks

import (
	"github.com/shimmeringbee/da"
	"github.com/stretchr/testify/mock"
)

var _ da.BasicCapability = (*BasicCapability)(nil)

type BasicCapability struct {
	mock.Mock
}

func (m *BasicCapability) Capability() da.Capability {
	args := m.Called()
	return args.Get(0).(da.Capability)
}

func (m *BasicCapability) Name() string {
	args := m.Called()
	return args.String(0)
}
