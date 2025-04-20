package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type RouteConfig struct {
	Name         string    `yaml:"name"`
	MatchPrefix  string    `yaml:"match_prefix"`
	Upstream     string    `yaml:"upstream"`
	AuthRequired bool      `yaml:"auth_required"`
	RateLimit    RateLimit `yaml:"rate_limit"`
}

type RateLimit struct {
	RPS   int `yaml:"rps"`
	Burst int `yaml:"burst"`
}

type Config struct {
	Routes     []RouteConfig       `yaml:"routes"`
	ApiKeys    []string            `yaml:"api_keys"`
	MapApiKeys map[string]struct{} `yaml:"-"`
}

func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	cfg.MapApiKeys = make(map[string]struct{}, len(cfg.ApiKeys))
	for _, key := range cfg.ApiKeys {
		cfg.MapApiKeys[key] = struct{}{}
	}

	return &cfg, nil
}
