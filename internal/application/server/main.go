package server

import (
	"context"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
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

func (i *Server) Start(ctx context.Context) error {
	ln, err := reuseport.Listen(
		i.listenNetwork,
		i.listenAddr)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		err := i.httpServer.Shutdown()
		if err != nil {
			i.logger.Errorf("server shutdown", "error", err.Error())
		}
	}()

	return i.httpServer.Serve(ln)
}
