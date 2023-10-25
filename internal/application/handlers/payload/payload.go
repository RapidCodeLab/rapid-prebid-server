package payload_handler

import (
	"encoding/json"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
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
		map[interfaces.EntityType][]interfaces.Entity,
	)

	for _, entityID := range req.Entities {
		entity, err := h.entityProvider.Provide(entityID)
		if err != nil {
			h.logger.Errorf("entity provider: %s", err.Error())
			continue
		}
		entities[entity.Type] = append(
			entities[entity.Type],
			entity,
		)
	}

	deviceData := h.deviceDetector.Detect(
		string(ctx.UserAgent()),
	)

	geoData, err := h.geoDetector.Detect()
	if err != nil {
		h.logger.Errorf("geo detect: %s", err.Error())
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
}
