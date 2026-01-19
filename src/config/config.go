package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	App struct {
		MaxFileSizeMB int64 `yaml:"max_file_size_mb"`
		ResolutionDPI int   `yaml:"resolution_dpi"`
	} `yaml:"app"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
