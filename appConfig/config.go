package appConfig

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Binlog struct {
		Directory           string `yaml:"directory"`
		Oversize            int64  `yaml:"oversize"`
		EveryCheckOversize  bool   `yaml:"every_check_oversize"`
		ChanceCheckOversize int    `yaml:"chance_check_oversize"`
	} `yaml:"binlog"`
	Snapshot struct {
		Directory string `yaml:"directory"`
	} `yaml:"snapshot"`
}

func LoadConfig(filename string) (*Config, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
