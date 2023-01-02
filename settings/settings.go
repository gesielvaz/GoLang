package settings

import (
	_ "embed"

	"gopkg.in/yaml.v2"
)

//go:embed settings.yaml
var settingsFile []byte

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}
type Settings struct {
	Port    string         `yaml:"port"`
	DB      DatabaseConfig `yaml:"database"`
	Version string         `yaml:"version"`
}

func New() *Settings {
	var s Settings
	err := yaml.Unmarshal(settingsFile, &s);
	if err != nil {
		return nil
	}
	return &s
}
