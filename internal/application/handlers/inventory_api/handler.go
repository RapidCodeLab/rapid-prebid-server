package inventoryapi_handler

import (
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
)

const (
	contentTypeApplicationJson = "application/json"
	idUserValue                = "id"
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

