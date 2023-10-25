package payload_handler

import (
	"github.com/buaazp/fasthttprouter"
)

func (h *Handler) LoadRoutes(r *fasthttprouter.Router) {
	r.POST("/rtb/v1", h.Handle)
}
