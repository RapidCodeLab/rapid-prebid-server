package core

import (
	"context"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
)

type payloadHTTPServer interface {
	Start(context.Context, interfaces.InventoryStorager) error
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
) error {
	return i.payloadHTTPServer.Start(ctx, invStorager)
}
