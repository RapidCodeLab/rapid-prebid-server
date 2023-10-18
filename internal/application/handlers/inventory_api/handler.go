package inventoryapi_handler

import (
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	logger       interfaces.Logger
	inventoryApi interfaces.InventoryAPI
	// possible data struct for healtchek response
}

func New(
	l interfaces.Logger,
	api interfaces.InventoryAPI,
) *Handler {
	return &Handler{
		logger:       l,
		inventoryApi: api,
	}
}

func (h *Handler) HandleReadAllInventories(ctx *fasthttp.RequestCtx) {
}

func (h *Handler) HandleCreateInventory(ctx *fasthttp.RequestCtx) {
}

func (h *Handler) HandleReadInventory(ctx *fasthttp.RequestCtx) {
}

func (h *Handler) HandleUpdateInventory(ctx *fasthttp.RequestCtx) {
}

func (h *Handler) HandleDeleteInventory(ctx *fasthttp.RequestCtx) {
}

func (h *Handler) HandleReadAllEntities(ctx *fasthttp.RequestCtx) {
}

func (h *Handler) HandleCreateEntity(ctx *fasthttp.RequestCtx) {
}

func (h *Handler) HandleReadEntity(ctx *fasthttp.RequestCtx) {
}

func (h *Handler) HandleUpdateEntity(ctx *fasthttp.RequestCtx) {
}

func (h *Handler) HandleDeleteEntity(ctx *fasthttp.RequestCtx) {
}
