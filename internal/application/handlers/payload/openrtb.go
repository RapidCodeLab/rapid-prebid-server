package payload_handler

import (
	"encoding/json"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/google/uuid"
	"github.com/prebid/openrtb/v17/adcom1"
	"github.com/prebid/openrtb/v17/native1"
	"github.com/prebid/openrtb/v17/native1/request"
	"github.com/prebid/openrtb/v17/openrtb2"
)

func initBidRequest(
	deviceData interfaces.DeviceData,
	geoData interfaces.GeoData,
	inventoryID string,
	inventoryType interfaces.InventoryType,
	inventoryIABCategories []string,
	inventoryIABCategoriesTaxonomy adcom1.CategoryTaxonomy,
) openrtb2.BidRequest {
	bidRequest := openrtb2.BidRequest{}

	bidRequest.ID = uuid.New().String()
	bidRequest.Device = &openrtb2.Device{
		UA:         deviceData.UserAgent,
		DeviceType: deviceData.DeviceType,
		OS:         deviceData.Platform,
		IP:         geoData.IP,
		IPv6:       geoData.IPv6,
	}
	bidRequest.Device.Geo = &openrtb2.Geo{
		Country: geoData.CountryCode,
		Region:  geoData.Region,
		City:    geoData.City,
	}

	switch inventoryType {
	case interfaces.InventoryTypeSite:
		bidRequest.Site = &openrtb2.Site{
			ID:     inventoryID,
			Cat:    inventoryIABCategories,
			CatTax: inventoryIABCategoriesTaxonomy,
		}

	case interfaces.InventoryTypeApp:
		bidRequest.App = &openrtb2.App{
			ID:     inventoryID,
			Cat:    inventoryIABCategories,
			CatTax: inventoryIABCategoriesTaxonomy,
		}
	}

	return bidRequest
}

func prepareImpObjects(
	entities []interfaces.Entity,
) []openrtb2.Imp {
	imps := []openrtb2.Imp{}

	for _, entity := range entities {
		imp := openrtb2.Imp{
			ID: entity.ID,
		}
		switch entity.Type {
		case interfaces.EntityTypeBanner:
			imp.Banner = &openrtb2.Banner{
				W: &entity.Width,
				H: &entity.Height,
			}
		case interfaces.EntityTypeNative:
			nativeRequest, err := prepareNativeObject(
				entity,
			)
			if err != nil {
				continue
			}
			imp.Native = &openrtb2.Native{
				Request: string(nativeRequest),
			}
		}
		imps = append(imps, imp)
	}
	return imps
}

func prepareNativeObject(
	entity interfaces.Entity,
) ([]byte, error) {
	data := request.Request{
		Ver:      "1.2",
		PlcmtCnt: entity.PlacementCount,
	}

	titleAsset := request.Asset{
		Required: 1,
		Title: &request.Title{
			Len: entity.TitleLength,
		},
	}

	imageAsset := request.Asset{
		Required: 1,
		Img: &request.Image{
			Type: native1.ImageAssetTypeMain,
			WMin: entity.Width,
			HMin: entity.Height,
		},
	}

	dataAsset := request.Asset{
		Required: 1,
		Data: &request.Data{
			Type: native1.DataAssetTypeDesc,
		},
	}

	data.Assets = append(
		data.Assets,
		titleAsset,
		imageAsset,
		dataAsset,
	)

	return json.Marshal(data)
}
