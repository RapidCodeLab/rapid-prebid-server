package inventoryapi_handler

import (
	"encoding/json"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/valyala/fasthttp"
)

func (h *Handler) HandleReadAllEntities(ctx *fasthttp.RequestCtx) {
	data, err := h.inventoryApi.ReadAllEntities()
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

func (h *Handler) HandleCreateEntity(ctx *fasthttp.RequestCtx) {
	entity := interfaces.Entity{}

	err := json.Unmarshal(ctx.PostBody(), &entity)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err = h.inventoryApi.CreateEntity(entity)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}
}

func (h *Handler) HandleReadEntity(ctx *fasthttp.RequestCtx) {
	ID := ctx.UserValue(idUserValue).(string)
	data, err := h.inventoryApi.ReadEntity(ID)
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

func (h *Handler) HandleUpdateEntity(ctx *fasthttp.RequestCtx) {
	entity := interfaces.Entity{}

	err := json.Unmarshal(ctx.PostBody(), &entity)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err = h.inventoryApi.UpdateEntity(entity)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}
}

func (h *Handler) HandleDeleteEntity(ctx *fasthttp.RequestCtx) {
	ID := ctx.UserValue(idUserValue).(string)
	err := h.inventoryApi.DeleteEntity(ID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}
}

