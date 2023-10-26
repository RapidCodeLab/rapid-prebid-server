package payload_handler

import (
	"encoding/json"
	"net"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
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

	deviceData := h.deviceDetector.Detect(
		string(ctx.UserAgent()),
	)

	geoData, err := h.geoDetector.Detect(
		net.ParseIP(ctx.RemoteIP().String()),
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
		entities[0].InventoryType,
		&bidRequest,
	)

	res := payloadResponse{}
	data, err := json.Marshal(res)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}

	ctx.SetContentType(contentTypeApplicationJson)
	ctx.SetBody(data)
}
