package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/core"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/server"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		cancel()
	}()
	payload_server := server.New()

	app := core.New(payload_server)

	err := app.Start(ctx)
	if err != nil {
		os.Exit(1)
	}
}
