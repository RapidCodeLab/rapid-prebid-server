package payload_handler

import (
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/google/uuid"
	"github.com/prebid/openrtb/v17/openrtb2"
)

func initBidRequest(
	deviceData interfaces.DeviceData,
	geoData interfaces.GeoData,
	inventoryType interfaces.InventoryType,
	bidRequest *openrtb2.BidRequest,
) {
	bidRequest.ID = uuid.New().String()
	bidRequest.Device = &openrtb2.Device{}
	bidRequest.Device.Geo = &openrtb2.Geo{}

	switch inventoryType {
	case interfaces.InventoryTypeSite:
		bidRequest.Site = &openrtb2.Site{}
	case interfaces.InventoryTypeApp:
		bidRequest.App = &openrtb2.App{}
	}
}
