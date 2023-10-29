package default_dsp_adapter

import (
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/prebid/openrtb/v17/openrtb2"
)

type adapter struct{}

func (i *adapter) DoRequest(
	req openrtb2.BidRequest,
) (
	openrtb2.BidResponse,
	error,
) {
	return openrtb2.BidResponse{}, nil
}

func NewDSPAdpater() (interfaces.DSPAdapter, error) {
	return &adapter{}, nil
}
