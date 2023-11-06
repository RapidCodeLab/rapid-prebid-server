package payload_handler

import (
	"encoding/json"
	"net"

	default_auction "github.com/RapidCodeLab/rapid-prebid-server/auctions/default"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	ipdetect "github.com/RapidCodeLab/rapid-prebid-server/pkg/ip-detect"
	"github.com/valyala/fasthttp"
)

func (h *Handler) Handle(ctx *fasthttp.RequestCtx) {
	req := payloadRequest{}

	err := json.Unmarshal(ctx.PostBody(), &req)
	if err != nil ||
		len(req.Entities) < 1 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	entities := make(
		[]interfaces.Entity, 0, len(req.Entities),
	)

	for _, entityID := range req.Entities {
		entity, err := h.entityProvider.Provide(entityID)
		if err != nil {
			h.logger.Errorf("entity provider: %s", err.Error())
			continue
		}
		entities = append(
			entities,
			entity,
		)
	}

	if len(entities) < 1 {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	deviceData := h.deviceDetector.Detect(
		string(ctx.UserAgent()),
	)

	geoData, err := h.geoDetector.Detect(
		net.ParseIP(
			ipdetect.FromRequest(ctx),
		),
	)
	if err != nil {
		h.logger.Errorf("geo detect: %s", err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bidRequest := initBidRequest(
		deviceData,
		geoData,
		entities[0].InventoryID,
		entities[0].InventoryType,
		entities[0].IABCategories,
		entities[0].IABCategoriesTaxonomy,
	)

	bidRequest.Imp = prepareImpObjects(entities)

	responses := h.doRequests(bidRequest)

	if len(responses) < 1 {
		ctx.SetStatusCode(fasthttp.StatusNoContent)
		return
	}

	placementCountMeta := make(map[string]int64)
	for _, e := range entities {
		placementCountMeta[e.ID] = e.PlacementCount
	}

	winners := default_auction.Auction(
		responses,
		placementCountMeta,
	)
	if len(winners) < 1 {
		ctx.SetStatusCode(fasthttp.StatusNoContent)
		return
	}

	res := buildResponse(winners)
	data, err := json.Marshal(res)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}

	ctx.SetContentType(contentTypeApplicationJson)
	ctx.SetBody(data)
}
