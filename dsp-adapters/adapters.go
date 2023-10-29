package dspadapters

import (
	default_dsp_adapter "github.com/RapidCodeLab/rapid-prebid-server/dsp-adapters/default"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
)

const (
	DefaultAdapter = "default"
)

func AdaptersInitializers() map[interfaces.DSPName]interfaces.NewDSPAdapter {
	return map[interfaces.DSPName]interfaces.NewDSPAdapter{
		DefaultAdapter: default_dsp_adapter.NewDSPAdpater,
	}
}
