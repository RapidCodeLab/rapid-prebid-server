package core

import (
	"context"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
)

type payloadHTTPServer interface {
	Start(context.Context, interfaces.InventoryAPI) error
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
	invAPI interfaces.InventoryAPI,
) error {
	return i.payloadHTTPServer.Start(ctx, invAPI)
}
