package interfaces

import (
	"github.com/prebid/openrtb/v17/openrtb2"
)

type (
	DSPName string

	DSPConfigProvider interface {
		Read(DSPName) (DSPAdapterConfig, error)
	}

	DSPAdapter interface {
		GetName() DSPName
		DoRequest(
			openrtb2.BidRequest) (
			openrtb2.BidResponse,
			error)
	}

	DSPAdapterConfig struct {
		Name          DSPName `json:"name"`
		Endpoint      string  `json:"endpoint"`
		RequestTimout int64   `json:"request_timeout"`
	}

	NewDSPAdapter func(DSPAdapterConfig) (DSPAdapter, error)
)
