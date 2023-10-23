package inventoryapi_handler

import (
	"encoding/json"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/valyala/fasthttp"
)

func (h *Handler) HandleReadAllInventories(ctx *fasthttp.RequestCtx) {
	data, err := h.inventoryApi.ReadAllInventories()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}
	if len(data) < 1 {
		ctx.SetStatusCode(fasthttp.StatusNoContent)
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}

	ctx.SetContentType(contentTypeApplicationJson)
	ctx.SetBody(res)
}

func (h *Handler) HandleCreateInventory(ctx *fasthttp.RequestCtx) {
	inv := interfaces.Inventory{}

	err := json.Unmarshal(ctx.PostBody(), &inv)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err = h.inventoryApi.CreateInventory(inv)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}
}

func (h *Handler) HandleReadInventory(ctx *fasthttp.RequestCtx) {
	ID := ctx.UserValue(idUserValue).(string)

	data, err := h.inventoryApi.ReadInventory(ID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}

	ctx.SetContentType(contentTypeApplicationJson)
	ctx.SetBody(res)
}

func (h *Handler) HandleUpdateInventory(ctx *fasthttp.RequestCtx) {
	inv := interfaces.Inventory{}

	err := json.Unmarshal(ctx.PostBody(), &inv)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err = h.inventoryApi.UpdateInventory(inv)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}
}

func (h *Handler) HandleDeleteInventory(ctx *fasthttp.RequestCtx) {
	ID := ctx.UserValue(idUserValue).(string)
	err := h.inventoryApi.DeleteInventory(ID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}
}
