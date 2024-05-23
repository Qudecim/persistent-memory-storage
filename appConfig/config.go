package appConfig

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Binlog struct {
		Directory string `yaml:"directory"`
		MaxWrites int    `yaml:"max_writes"`
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
