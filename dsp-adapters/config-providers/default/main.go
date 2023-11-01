package default_config_provider

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/RapidCodeLab/rapid-prebid-server/internal/application/interfaces"
)

var NoConfigFilesErr = errors.New("no config files found in path")

type provider struct {
	data map[interfaces.DSPName]interfaces.DSPAdapterConfig
}

func New(
	path string,
	l interfaces.Logger,
) (
	interfaces.DSPConfigProvider,
	error,
) {
	configFiles, err := findFiles(path, ".json")
	if err != nil {
		return nil, err
	}
	if len(configFiles) < 1 {
		return nil, NoConfigFilesErr
	}

	provider := &provider{}
	provider.data = make(map[interfaces.DSPName]interfaces.DSPAdapterConfig)

	for _, filePath := range configFiles {
		f, err := os.Open(filePath)
		if err != nil {
			l.Errorf("read dsp config file: %s", err.Error())
			continue
		}

		decoder := json.NewDecoder(f)
		config := interfaces.DSPAdapterConfig{}
		err = decoder.Decode(&config)
		if err != nil {
			l.Errorf("decode config:%s", err.Error())
			continue
		}
		provider.data[config.Name] = config
	}

	return provider, nil
}

func (i *provider) Read(
	name interfaces.DSPName,
) (
	interfaces.DSPAdapterConfig,
	error,
) {
	return i.data[name], nil
}

func findFiles(root, ext string) ([]string, error) {
	var a []string
	err := filepath.WalkDir(
		root,
		func(s string, d fs.DirEntry, e error) error {
			if e != nil {
				return e
			}
			if filepath.Ext(d.Name()) == ext {
				a = append(a, s)
			}
			return nil
		})
	return a, err
}
