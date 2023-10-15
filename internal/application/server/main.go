package server

import "github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"

type Server struct{
	logger interfaces.Logger
}

func New() *Server {
	return &Server{}
}

func Start() error {
	return nil
}
