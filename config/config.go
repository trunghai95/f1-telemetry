package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// AppConfig stores the necessary configs, duh
type AppConfig struct {
	UDPListen string
}

var cfg *AppConfig

// GetAppConfig returns the app config
// If uninitialized, this returns nil
func GetAppConfig() *AppConfig {
	return cfg
}

// InitConfigYaml initializes the config using the file at configPath
// This reads the config file as a Yaml file
func InitConfigYaml(configPath string) (err error) {
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return
	}
	if cfg == nil {
		cfg = new(AppConfig)
	}

	err = yaml.Unmarshal(b, cfg)
	return
}
