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
	browscap_devicedetector "github.com/RapidCodeLab/rapid-prebid-server/device-detectors/browscap"
	default_config_provider "github.com/RapidCodeLab/rapid-prebid-server/dsp-adapters/config-providers/default"
	geoip2_detector "github.com/RapidCodeLab/rapid-prebid-server/geo-detectors/geoip2"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/core"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/server"
	inventoryapiboltdb "github.com/RapidCodeLab/rapid-prebid-server/inventory-api/boltdb"
)

type Config struct {
	Debug                     bool   `env:"DEBUG"`
	ServerListeNetwork        string `env:"SERVER_LISTEN_NETWORK,required"`
	ServerListenAddr          string `env:"SERVER_LISTEN_ADDR,required"`
	BoltDBPath                string `env:"BOLTDB_PATH,required"`
	DeviceDBPath              string `env:"DEVICE_DB_PATH,required"`
	GeoDBPath                 string `env:"GEO_DB_PATH,required"`
	DSPAdaptersConfigFilePath string `env:"DSP_ADPATERS_CONFIG_FILE_PATH, required"`
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

	l.Infof("application started. %s\n", time.Now().Format("2006-01-02 15:04:05"))

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
		l.Errorf("boltDB open: %s\n", err.Error())
		os.Exit(1)
	}

	deviceDetector, err := browscap_devicedetector.New(config.DeviceDBPath)
	if err != nil {
		l.Errorf("deviceDB open: %s\n", err.Error())
		os.Exit(1)
	}

	geoDetector, err := geoip2_detector.New(config.GeoDBPath, "ru")
	if err != nil {
		l.Errorf("geoDB open: %s\n", err.Error())
		os.Exit(1)
	}

	app := core.New(s, l)

	enabledDSPAdapters := []interfaces.DSPName{}
	dspConfigProvider, err := default_config_provider.New(config.DSPAdaptersConfigFilePath)
	if err != nil {
		l.Errorf("dsp config provider: %s\n", err.Error())
		os.Exit(1)
	}
	err = app.Start(
		ctx,
		invStorager,
		deviceDetector,
		geoDetector,
		enabledDSPAdapters,
		dspConfigProvider,
	)
	if err != nil {
		l.Errorf("app exit", err.Error())
		os.Exit(1)
	}

	l.Infof("application successfully stopped. %s", time.Now().Format("2006-01-02 15:04:05"))
}
