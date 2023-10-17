package payload_handler

import (
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/valyala/fasthttp"
)

type payloadRequest struct {
	Inventories []string `json:"inventories"`
}

type payloadResponse struct {
	Paylads []payload `json:"payloads"`
}

type payload struct {
	InventoryID string `json:"inventory_id"`
	Type        string `json:"type"`
	Adm         string `json:"adm"`
}

type Handler struct {
	logger interfaces.Logger
	// possible data struct for healtchek response
}

func New(l interfaces.Logger) *Handler {
	return &Handler{
		logger: l,
	}
}

func (h *Handler) Handle(ctx *fasthttp.RequestCtx) {
}
