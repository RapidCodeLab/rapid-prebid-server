package server

import (
	"context"
	"errors"

	dspadapters "github.com/RapidCodeLab/rapid-prebid-server/dsp-adapters"
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

func New(l interfaces.Logger,
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

	adaptersInitializers := dspadapters.AdaptersInitializers()
	adapters := []interfaces.DSPAdapter{}

	for _, dspName := range enabledAdapters {
		config, err := dspConfigProvider.Read(dspName)
		if err != nil {
			i.logger.Errorf("dsp adapter config read: %s", err.Error())
			continue
		}

		init, ok := adaptersInitializers[dspName]
		if !ok {
			i.logger.Errorf("dsp initializer not found: %s", dspName)
			continue
		}

		adapter, err := init(config)
		if err != nil {
			i.logger.Errorf("dsp adapter init: %s", err.Error())
			continue
		}
		adapters = append(adapters, adapter)
	}

	if len(adapters) < 1 {
		return NoEnabledDSPAdaptersErr
	}

	payloadHandler := payload_handler.New(
		i.logger,
		deviceDetector,
		geoDetector,
		nil,
		adapters,
	)
	payloadHandler.LoadRoutes(r)

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
