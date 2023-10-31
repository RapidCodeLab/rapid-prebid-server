package payload_handler

import (
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/prebid/openrtb/v17/openrtb2"
)

const (
	contentTypeApplicationJson = "application/json"
)

type payloadRequest struct {
	Entities []string `json:"entities"`
}

type payloadResponse struct {
	Paylads []payload `json:"payloads"`
}

type payload struct {
	EntityID   string              `json:"entity_id"`
	MarkupType openrtb2.MarkupType `json:"markup_type"`
	Adm        string              `json:"adm"`
}

type Handler struct {
	logger         interfaces.Logger
	deviceDetector interfaces.DeviceDetector
	geoDetector    interfaces.GeoDetector
	entityProvider interfaces.EntityProvider
	dspAdapters    []interfaces.DSPAdapter
	// possible data struct for healtchek response
}

func New(
	l interfaces.Logger,
	dd interfaces.DeviceDetector,
	gd interfaces.GeoDetector,
	p interfaces.EntityProvider,
	a []interfaces.DSPAdapter,
) *Handler {
	return &Handler{
		logger:         l,
		deviceDetector: dd,
		geoDetector:    gd,
		entityProvider: p,
		dspAdapters:    a,
	}
}
