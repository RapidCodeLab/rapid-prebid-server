package inventoryapi_handler

import (
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
)

const (
	contentTypeApplicationJson = "application/json"
	idUserValue                = "id"
)

type Handler struct {
	logger      interfaces.Logger
	invStorager interfaces.InventoryStorager
	// possible data struct for healtchek response
}

func New(
	l interfaces.Logger,
	s interfaces.InventoryStorager,
) *Handler {
	return &Handler{
		logger:      l,
		invStorager: s,
	}
}
