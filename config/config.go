package config

import (
	"fmt"
	"os"
	"time"

	"github.com/supergoodsystems/supergood-proxy/proxy"
	"github.com/supergoodsystems/supergood-proxy/remoteconfigworker"
)

type Config struct {
	RemoteWorkerConfig remoteconfigworker.RemoteConfigOpts `yaml:"remoteworkerConfig"`
	ProxyConfig        proxy.ProxyOpts `yaml:"proxyConfig"`
}

func GetConfig(path string) Config {
	cfg := Config{}
	resolveConfigWithPath(path, &cfg)
	resolveConfigWithEnv(&cfg)
	return cfg
}

func resolveConfigWithEnv(config *Config) error {
	if (config.RemoteWorkerConfig.BaseURL == "") {
		config.RemoteWorkerConfig.BaseURL = os.Getenv("SUPERGOOD_BASE_URL")
		if config.RemoteWorkerConfig.BaseURL == "" {
			config.RemoteWorkerConfig.BaseURL = "http://localhost:3001"
		}
	}

	if (config.RemoteWorkerConfig.AdminClientID == "") {
		config.RemoteWorkerConfig.AdminClientID = os.Getenv("ADMIN_CLIENT_ID")
		if config.RemoteWorkerConfig.AdminClientID == "" {
			return fmt.Errorf("ADMIN_CLIENT_ID missing from env vars")
		}
	}

	if (config.RemoteWorkerConfig.AdminClientSecret == "") {
		config.RemoteWorkerConfig.AdminClientSecret = os.Getenv("ADMIN_CLIENT_SECRET")
		if config.RemoteWorkerConfig.AdminClientSecret == "" {
			return fmt.Errorf("ADMIN_CLIENT_SECRET missing from env vars")
		}
	}

	if (config.RemoteWorkerConfig.FetchInterval == 0) {
		config.RemoteWorkerConfig.FetchInterval = 1 * time.Second
	}

	if (config.ProxyConfig.Port == "") {
		config.ProxyConfig.Port = os.Getenv("PROXY_HTTP_PORT")
		if config.ProxyConfig.Port  == "" {
			config.ProxyConfig.Port  = "8080"
		}
	}

	return nil
}
