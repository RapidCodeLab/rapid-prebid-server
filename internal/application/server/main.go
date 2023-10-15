package server

import (
	"context"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/valyala/fasthttp"
)

type Server struct {
	logger     interfaces.Logger
	httpServer *fasthttp.Server
}

func New(l interfaces.Logger) *Server {
	return &Server{
		logger: l,
		httpServer: &fasthttp.Server{
		},
	}
}

func (i *Server) Start(ctx context.Context) error {
	go func() {
		<-ctx.Done()
	}()
	return i.httpServer.Serve(nil)
}
