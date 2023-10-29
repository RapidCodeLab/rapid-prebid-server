package interfaces

import (
	"github.com/prebid/openrtb/v17/openrtb2"
)

type DSPName string

type DSPAdapter interface {
	DoRequest(
		openrtb2.BidRequest) (
		openrtb2.BidResponse,
		error)
}

type DSPAdapterConfig struct {
	Name     DSPName
	Endpoint string
}

type NewDSPAdapter func() (DSPAdapter, error)
