package interfaces

import (
	"context"

	"github.com/prebid/openrtb/v17/openrtb2"
)

type DSPAdapter interface{
   DoRequest(
		 context.Context,
		 openrtb2.BidRequest) (
			 openrtb2.BidResponse,
			 error)
}
