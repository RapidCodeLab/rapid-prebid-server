package default_config_provider

import "github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"

type provider struct{}

func New(
	path string,
) (
	interfaces.DSPConfigProvider,
	error,
) {
	return &provider{}, nil
}

func (i *provider) Read(
	name interfaces.DSPName,
) (
	interfaces.DSPAdapterConfig,
	error,
) {
	return interfaces.DSPAdapterConfig{}, nil
}
