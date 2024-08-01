package config

import (
	"log"
	"os"
	"time"

	"github.com/supergoodsystems/supergood-proxy/proxy"
	"github.com/supergoodsystems/supergood-proxy/remoteconfigworker"
)

type Config struct {
	RemoteWorkerConfig remoteconfigworker.RemoteConfigOpts
	ProxyConfig proxy.ProxyOpts
}

func GetConfigFromEnvironment() Config {
	baseURL := os.Getenv("SUPERGOOD_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:3001"
	}
	adminClientId := os.Getenv("ADMIN_CLIENT_ID")
	if adminClientId == "" {
		log.Fatal("ADMIN_CLIENT_ID missing from env vars")
	}

	adminClientSecret := os.Getenv("ADMIN_CLIENT_SECRET")
	if adminClientSecret == "" {
		log.Fatal("ADMIN_CLIENT_SECRET missing from env vars")
	}

	fetchInterval := 1 * time.Second

	proxyPort := os.Getenv("PROXY_HTTP_PORT")
	if proxyPort == ""{
		proxyPort = "8080"
	}

	return Config{
		RemoteWorkerConfig: remoteconfigworker.RemoteConfigOpts{
			BaseURL:           baseURL,
			AdminClientID:     adminClientId,
			AdminClientSecret: adminClientSecret,
			FetchInterval:     fetchInterval,
		},
		ProxyConfig: proxy.ProxyOpts{
			Port: proxyPort,
		},
	}

}
