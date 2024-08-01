package remoteconfigworker

import (
	"os"
	"time"
)

// GetConfigFromEvironment returns the proxy config pulled from environment variables
func GetConfigFromEnvironment() RemoteConfigOpts {
	baseURL := os.Getenv("SUPERGOOD_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:3001"
	}
	adminClientId := os.Getenv("ADMIN_CLIENT_ID")
	if adminClientId == "" {
		adminClientId = ""
	}

	adminClientSecret := os.Getenv("ADMIN_CLIENT_SECRET")
	if adminClientSecret == "" {
		adminClientSecret = ""
	}
	fetchInterval := 1 * time.Second

	return RemoteConfigOpts{
		BaseURL:           baseURL,
		AdminClientID:     adminClientId,
		AdminClientSecret: adminClientSecret,
		FetchInterval:     fetchInterval,
	}
}
