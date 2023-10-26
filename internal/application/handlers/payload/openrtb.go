package payload_handler

import (
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/prebid/openrtb/v17/openrtb2"
)

func initBidRequest(
	deviceData interfaces.DeviceData,
	geoData interfaces.GeoData,
	inventoryType interfaces.InventoryType,
	bidRequest *openrtb2.BidRequest,
) {
}
