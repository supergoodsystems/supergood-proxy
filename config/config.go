package config

import (
	"fmt"
	"os"
	"time"

	"github.com/supergoodsystems/supergood-proxy/proxy"
	"github.com/supergoodsystems/supergood-proxy/remoteconfigworker"
)

type Config struct {
	RemoteWorkerConfig remoteconfigworker.RemoteConfigOpts `yaml:"remoteWorkerConfig"`
	ProxyConfig        proxy.ProxyOpts                     `yaml:"proxyConfig"`
}

func GetConfig(path string) (Config, error) {
	cfg := Config{}
	var err error
	if path == "" {
		if path, err = resolvePathFromEnv(); err != nil {
			return cfg, err
		}
	}

	resolveConfigWithPath(path, &cfg)
	err = resolveConfigWithEnv(&cfg)
	return cfg, err
}

func resolvePathFromEnv() (string, error) {
	env := os.Getenv("ENV")
	if env == "" {
		return "", fmt.Errorf("cannot have env path undefined as well as ENV var undefined")
	}
	if env == "development" {
		return "_config/dev.yml", nil
	}
	if env == "staging" {
		return "/var/_config/staging.yml", nil
	}
	if env == "production" {
		return "/var/_config/production.yml", nil
	}
	return "", fmt.Errorf("cannot resolve path from environment. Invalid ENV var")
}

func resolveConfigWithEnv(config *Config) error {
	if config.RemoteWorkerConfig.BaseURL == "" {
		config.RemoteWorkerConfig.BaseURL = os.Getenv("SUPERGOOD_BASE_URL")
		if config.RemoteWorkerConfig.BaseURL == "" {
			config.RemoteWorkerConfig.BaseURL = "http://localhost:3001"
		}
	}

	if config.RemoteWorkerConfig.AdminClientKey == "" {
		config.RemoteWorkerConfig.AdminClientKey = os.Getenv("ADMIN_CLIENT_KEY")
		if config.RemoteWorkerConfig.AdminClientKey == "" {
			return fmt.Errorf("ADMIN_CLIENT_KEY missing from env vars")
		}
	}

	if config.RemoteWorkerConfig.FetchInterval == 0 {
		config.RemoteWorkerConfig.FetchInterval = 1 * time.Second
	}

	if config.ProxyConfig.Port == "" {
		config.ProxyConfig.Port = os.Getenv("PROXY_HTTP_PORT")
		if config.ProxyConfig.Port == "" {
			config.ProxyConfig.Port = "8080"
		}
	}

	return nil
}
