package dspadapters

import (
	default_dsp_adapter "github.com/RapidCodeLab/rapid-prebid-server/dsp-adapters/default"
	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
)

const (
	DemoDSPAdapter1 = "demo-dsp-1"
	DemoDSPAdapter2 = "demo-dsp-2"
)

func AdaptersInitializers() map[interfaces.DSPName]interfaces.NewDSPAdapter {
	return map[interfaces.DSPName]interfaces.NewDSPAdapter{
		DemoDSPAdapter1: default_dsp_adapter.NewDSPAdpater,
		DemoDSPAdapter2: default_dsp_adapter.NewDSPAdpater,
	}
}
