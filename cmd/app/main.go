package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env"

	zaplogger "github.com/RapidCodeLab/ZapLogger"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/core"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/server"
)

type Config struct {
	Debug              bool   `env:"DEBUG"`
	ServerListeNetwork string `env:"SERVER_LISTEN_NETWORK,required"`
	ServerListenAddr   string `env:"SERVER_LISTEN_ADDR,required"`
}

func main() {
	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		fmt.Printf("config parse error: %s", err.Error())
		os.Exit(1)
	}

	l, err := zaplogger.New(&zaplogger.Config{})
	if err != nil {
		os.Exit(1)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		cancel()
	}()

	s := server.New(l,
		config.ServerListeNetwork,
		config.ServerListenAddr)

	app := core.New(s, l)

	err = app.Start(ctx)
	if err != nil {
		os.Exit(1)
	}
}
