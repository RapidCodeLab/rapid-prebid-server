package default_auction

import "github.com/prebid/openrtb/v17/openrtb2"

func Auction(
	responses []openrtb2.BidResponse,
) map[string][]openrtb2.Bid {
	auctionParticipants := make(
		map[string][]openrtb2.Bid,
	)

	for _, res := range responses {
		for _, seatBid := range res.SeatBid {
			for _, bid := range seatBid.Bid {
				auctionParticipants[bid.ImpID] = append(
					auctionParticipants[bid.ImpID],
					bid,
				)
			}
		}
	}

	winners := make(map[string][]openrtb2.Bid)
	return winners
}
