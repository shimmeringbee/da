package mocks

import (
	"context"
	"github.com/shimmeringbee/da"
	"github.com/shimmeringbee/da/capabilities"
	"github.com/stretchr/testify/mock"
)

var _ capabilities.ProductInformation = (*ProductInformation)(nil)

type ProductInformation struct {
	mock.Mock
}

func (m *ProductInformation) Get(c context.Context) (capabilities.ProductInfo, error) {
	args := m.Called(c)
	return args.Get(0).(capabilities.ProductInfo), args.Error(1)
}

func (m *ProductInformation) Capability() da.Capability {
	args := m.Called()
	return args.Get(0).(da.Capability)
}

func (m *ProductInformation) Name() string {
	args := m.Called()
	return args.String(0)
}
