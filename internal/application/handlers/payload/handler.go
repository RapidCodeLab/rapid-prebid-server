package payload_handler

import (
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/valyala/fasthttp"
)

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
