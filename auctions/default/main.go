package default_auction

import (
	"github.com/RapidCodeLab/rapid-prebid-server/auctions/algorythms/timsort"
	"github.com/prebid/openrtb/v17/openrtb2"
)

func ByPrice(a, b interface{}) bool {
	return a.(openrtb2.Bid).Price > b.(openrtb2.Bid).Price
}

// timsort algorytm
func Sort(participants []interface{}, seat int) []interface{} {
	timsort.Sort(participants, ByPrice)
	return participants[0:seat]
}

func Auction(
	responses []openrtb2.BidResponse,
) []interface{} {
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
		[]interface{},
		0,
		len(auctionParticipants),
	)

	for _, impParticipants := range auctionParticipants {
		impWinners := Sort(impParticipants, 1)
		winners = append(winners, impWinners...)
	}

	return winners
}
