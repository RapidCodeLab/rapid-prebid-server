package server

import (
	"context"
	"errors"

	inventoryapi_handler "github.com/RapidCodeLab/rapid-prebid-server/internal/application/handlers/inventory-api"
	payload_handler "github.com/RapidCodeLab/rapid-prebid-server/internal/application/handlers/payload"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
)

var NoEnabledDSPAdaptersErr = errors.New("At least one DSPAdapter must be enabled")

type Server struct {
	logger     interfaces.Logger
	httpServer *fasthttp.Server
	listenNetwork,
	listenAddr string
}

func New(
	l interfaces.Logger,
	network, addr string,
) *Server {
	return &Server{
		logger: l,
		httpServer: &fasthttp.Server{
			Name: "Rapid Prebid Server",
		},
		listenNetwork: network,
		listenAddr:    addr,
	}
}

func (i *Server) Start(
	ctx context.Context,
	invStorager interfaces.InventoryStorager,
	entityProvider interfaces.EntityProvider,
	deviceDetector interfaces.DeviceDetector,
	geoDetector interfaces.GeoDetector,
	enabledAdapters []interfaces.DSPName,
	dspConfigProvider interfaces.DSPConfigProvider,
) error {
	if len(enabledAdapters) < 1 {
		return NoEnabledDSPAdaptersErr
	}

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

	adapters, err := i.initDSPAdapters(
		enabledAdapters,
		dspConfigProvider,
	)
	if err != nil {
		return err
	}

	payloadHandler := payload_handler.New(
		i.logger,
		deviceDetector,
		geoDetector,
		entityProvider,
		adapters,
	)
	payloadHandler.LoadRoutes(r)

	i.httpServer.Handler = r.Handler

	go func() {
		<-ctx.Done()
		err := i.httpServer.Shutdown()
		if err != nil {
			i.logger.Errorf("server shutdown: %s", err.Error())
		}
	}()

	return i.httpServer.Serve(ln)
}
