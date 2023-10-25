package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env"

	zaplogger "github.com/RapidCodeLab/ZapLogger"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/core"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/server"
	inventoryapiboltdb "github.com/RapidCodeLab/rapid-prebid-server/inventory_api/boltdb"
)

type Config struct {
	Debug              bool   `env:"DEBUG"`
	ServerListeNetwork string `env:"SERVER_LISTEN_NETWORK,required"`
	ServerListenAddr   string `env:"SERVER_LISTEN_ADDR,required"`
	BoltDBPath         string `env:"BOLTDB_PATH,required"`
}

func main() {
	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		fmt.Printf("config parse error: %s\n", err.Error())
		os.Exit(1)
	}

	l, err := zaplogger.New(&zaplogger.Config{
		ServiceID:           "prebid",
		StdOutLoggerEnabled: true,
	})
	if err != nil {
		fmt.Printf("logger init: %s\n", err.Error())
		os.Exit(1)
	}

	l.Infof("application started. %s", time.Now().Format("2006-01-02 15:04:05"))

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

	invStorager, err := inventoryapiboltdb.New(
		ctx,
		config.BoltDBPath,
		l)
	if err != nil {
		l.Errorf("boltDB open", "err", err.Error())
		os.Exit(1)
	}

	app := core.New(s, l)

	err = app.Start(ctx, invStorager)
	if err != nil {
		l.Errorf("app exit", "err", err.Error())
		os.Exit(1)
	}

	l.Infof("application successfully stopped. %s", time.Now().Format("2006-01-02 15:04:05"))
}
