package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	MaxFileSizeMB int64 `yaml:"max_file_size_mb"`
	ResolutionDPI int   `yaml:"resolution_dpi"`
}

func LoadConfig(path string) (*Config, error) {

	defaultConfig := Config{
		MaxFileSizeMB: 10,
		ResolutionDPI: 300,
	}

	f, err := os.ReadFile(path)
	if err != nil {
		return &defaultConfig, err
	}

	var cfg Config
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return &defaultConfig, err
	}

	return &cfg, nil
}
