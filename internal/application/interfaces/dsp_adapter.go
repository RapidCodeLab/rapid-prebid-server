package interfaces

import (
	"github.com/prebid/openrtb/v17/openrtb2"
)

type DSPName string

type DSPConfigProvider interface {
	Read(DSPName) (DSPAdapterConfig, error)
}

type DSPAdapter interface {
	DoRequest(
		openrtb2.BidRequest) (
		openrtb2.BidResponse,
		error)
}

type DSPAdapterConfig struct {
	Name     DSPName `json:"name"`
	Endpoint string  `json:"endpoint"`
}

type NewDSPAdapter func(DSPAdapterConfig) (DSPAdapter, error)
