package capabilities

import (
	"context"
)

type ProductInfo struct {
	Manufacturer string
	Name         string
	Serial       string
	Version      string
}

type ProductInformation interface {
	Get(context.Context) (ProductInfo, error)
}
