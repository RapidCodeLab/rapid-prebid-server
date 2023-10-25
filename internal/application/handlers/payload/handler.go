package payload_handler

import (
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
)

type payloadRequest struct {
	Entities []string `json:"entities"`
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
	logger         interfaces.Logger
	deviceDetector interfaces.DeviceDetector
	geoDetector    interfaces.GeoDetector
	entityProvider interfaces.EntityProvider
	// possible data struct for healtchek response
}

func New(
	l interfaces.Logger,
	dd interfaces.DeviceDetector,
	gd interfaces.GeoDetector,
	p interfaces.EntityProvider,
) *Handler {
	return &Handler{
		logger:         l,
		deviceDetector: dd,
		geoDetector:    gd,
		entityProvider: p,
	}
}
