package payload_handler

import "github.com/prebid/openrtb/v17/openrtb2"

func buildResponse(
	bids []openrtb2.Bid,
) payloadResponse {
	res := payloadResponse{}

	for _, bid := range bids {
		p := payload{
			EntityID:   bid.ImpID,
			Adm:        bid.AdM,
			MarkupType: bid.MType,
		}
		res.Paylads = append(res.Paylads, p)
	}

	return res
}
