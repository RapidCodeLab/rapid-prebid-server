package health_handler

import (
	"encoding/json"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/valyala/fasthttp"
)

const (
	StatusOK = "OK"
)

type Response struct {
	Status string
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
	res := Response{
		Status: StatusOK,
	}
	body, err := json.Marshal(res)
	if err != nil {
		h.logger.Warnf("healtchek handler", "err", err)
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}
	ctx.SetBody(body)
	ctx.SetContentType("application/json")
}
