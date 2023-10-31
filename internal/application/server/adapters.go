package server

import (
	dspadapters "github.com/RapidCodeLab/rapid-prebid-server/dsp-adapters"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
)

func (i *Server) initDSPAdapters(
	enabledAdapters []interfaces.DSPName,
	dspConfigProvider interfaces.DSPConfigProvider,
) (
	[]interfaces.DSPAdapter,
	error,
) {
	adapters := []interfaces.DSPAdapter{}
	adaptersInitializers := dspadapters.AdaptersInitializers()
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
		return nil, NoEnabledDSPAdaptersErr
	}

	return adapters, nil
}
