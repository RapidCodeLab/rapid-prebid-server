package server

import (
	"context"

	inventoryapi_handler "github.com/RapidCodeLab/rapid-prebid-server/internal/application/handlers/inventory_api"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
)

type Server struct {
	logger     interfaces.Logger
	httpServer *fasthttp.Server
	listenNetwork,
	listenAddr string
}

func New(l interfaces.Logger,
	network, addr string,
) *Server {
	return &Server{
		logger:        l,
		httpServer:    &fasthttp.Server{},
		listenNetwork: network,
		listenAddr:    addr,
	}
}

func (i *Server) Start(ctx context.Context,
	invStorager interfaces.InventoryStorager,
) error {
	ln, err := reuseport.Listen(
		i.listenNetwork,
		i.listenAddr)
	if err != nil {
		return err
	}

	r := fasthttprouter.New()

	invAPIHandler := inventoryapi_handler.New(
		i.logger,
		invStorager,
	)
	invAPIHandler.LoadRoutes(r)

	i.httpServer.Handler = r.Handler

	go func() {
		<-ctx.Done()
		err := i.httpServer.Shutdown()
		if err != nil {
			i.logger.Errorf("server shutdown", "error", err.Error())
		}
	}()

	return i.httpServer.Serve(ln)
}
