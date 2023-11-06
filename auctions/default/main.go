package default_auction

import (
	"github.com/RapidCodeLab/rapid-prebid-server/auctions/algorythms/timsort"
	"github.com/prebid/openrtb/v17/openrtb2"
)

func ByPrice(a, b interface{}) bool {
	return a.(openrtb2.Bid).Price > b.(openrtb2.Bid).Price
}

// timsort algorytm
func Sort(
	participants []interface{},
	plcmntCount int64,
) []interface{} {
	timsort.Sort(participants, ByPrice)
	return participants[0:plcmntCount]
}

func Auction(
	responses []openrtb2.BidResponse,
	placementCountMeta map[string]int64,
) []openrtb2.Bid {
	auctionParticipants := make(
		map[string][]interface{},
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

	winners := make(
		[]openrtb2.Bid,
		0,
		len(auctionParticipants),
	)

	for entityID, impParticipants := range auctionParticipants {
		impWinners := Sort(
			impParticipants,
			placementCountMeta[entityID],
		)
		for _, w := range impWinners {
			winners = append(winners, w.(openrtb2.Bid))
		}
	}

	return winners
}
