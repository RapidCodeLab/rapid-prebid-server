package interfaces

import (
	"github.com/prebid/openrtb/v17/openrtb2"
)

type DSPAdapter interface {
	DoRequest(
		openrtb2.BidRequest) (
		openrtb2.BidResponse,
		error)
}
