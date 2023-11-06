package payload_handler

import (
	"sync"

	"github.com/prebid/openrtb/v17/openrtb2"
)

func (h *Handler) doRequests(
	bidRequest openrtb2.BidRequest,
) []openrtb2.BidResponse {
	responses := make(
		[]openrtb2.BidResponse,
		0,
		len(h.dspAdapters),
	)

	// request DSPs
	wg := sync.WaitGroup{}
	wg.Add(len(h.dspAdapters))

	for _, adapter := range h.dspAdapters {
		a := adapter
		go func() {
			defer wg.Done()
			bidResponse, err := a.DoRequest(bidRequest)
			if err != nil {
				h.logger.Errorf(
					"adapter %s request: %s",
					a.GetName(),
					err.Error(),
				)
				return
			}
			responses = append(responses, bidResponse)
		}()
	}

	wg.Wait()

	return responses
}
