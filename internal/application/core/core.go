package core

import (
	"context"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
)

type payloadHTTPServer interface {
	Start(
		context.Context,
		interfaces.InventoryStorager,
		interfaces.EntityProvider,
		interfaces.DeviceDetector,
		interfaces.GeoDetector,
		[]interfaces.DSPName,
		interfaces.DSPConfigProvider,
	) error
}

type Core struct {
	payloadHTTPServer payloadHTTPServer
	logger            interfaces.Logger
}

func New(s payloadHTTPServer,
	l interfaces.Logger,
) *Core {
	return &Core{
		payloadHTTPServer: s,
		logger:            l,
	}
}

func (i *Core) Start(ctx context.Context,
	invStorager interfaces.InventoryStorager,
	entityProvider interfaces.EntityProvider,
	deviceDetector interfaces.DeviceDetector,
	geoDetector interfaces.GeoDetector,
	enabledDSPAdapters []interfaces.DSPName,
	dspConfigProvider interfaces.DSPConfigProvider,
) error {
	return i.payloadHTTPServer.Start(
		ctx,
		invStorager,
		entityProvider,
		deviceDetector,
		geoDetector,
		enabledDSPAdapters,
		dspConfigProvider,
	)
}
