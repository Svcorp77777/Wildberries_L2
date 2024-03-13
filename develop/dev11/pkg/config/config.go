package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	Port string `yaml:"port"`
}

func (c *config) ReadСonfigurationFile(nameFile string) (config config, err error) {
	file, err := os.ReadFile(nameFile)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, nil
	}

	return config, nil
}

func Config() config {
	return config{}
}
