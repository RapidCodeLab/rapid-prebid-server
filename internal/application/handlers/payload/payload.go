package payload_handler

import (
	"encoding/json"
	"net"
	"sync"

	default_auction "github.com/RapidCodeLab/rapid-prebid-server/auctions/default"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	ipdetect "github.com/RapidCodeLab/rapid-prebid-server/pkg/ip-detect"
	"github.com/prebid/openrtb/v17/openrtb2"
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

	bidRequest := openrtb2.BidRequest{}

	initBidRequest(
		deviceData,
		geoData,
		entities[0].InventoryID,
		entities[0].InventoryType,
		entities[0].IABCategories,
		entities[0].IABCategoriesTaxonomy,
		&bidRequest,
	)

	bidRequest.Imp = prepareImpObjects(entities)

	responses := make([]openrtb2.BidResponse, 0, len(h.dspAdapters))

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
	if len(responses) < 1 {
		ctx.SetStatusCode(fasthttp.StatusNoContent)
		return
	}

	winners := default_auction.Auction(responses)
	if len(winners) < 1 {
		ctx.SetStatusCode(fasthttp.StatusNoContent)
		return
	}

	res := payloadResponse{}

	for _, winner := range winners {
		p := payload{
			EntityID:   winner.(openrtb2.Bid).ImpID,
			Adm:        winner.(openrtb2.Bid).AdM,
			MarkupType: winner.(openrtb2.Bid).MType,
		}
		res.Paylads = append(res.Paylads, p)
	}

	data, err := json.Marshal(res)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}

	ctx.SetContentType(contentTypeApplicationJson)
	ctx.SetBody(data)
}
