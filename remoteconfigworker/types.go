package remoteconfigworker

import (
	"time"

	"github.com/supergoodsystems/supergood-proxy/cache"
)

type RemoteConfigOpts struct {
	AdminClientID     string `yaml:"adminClientId"`
	AdminClientSecret string `yaml:"adminClientSecret"`
	BaseURL           string `yaml:"baseURL"`
	FetchInterval     time.Duration `yaml:"fetchInterval"`
}

type Credential struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type VendorConfig struct {
	Credentials []Credential `json:"credentials"`
}

type TenantConfig struct {
	ClientID     string                  `json:"clientID"`
	ClientSecret string                  `json:"clientSecret"`
	Vendors      map[string]VendorConfig `json:"vendors"` // map of domain to vendor config
}

type RemoteConfigWorker struct {
	adminClientId     string
	adminClientSecret string
	baseURL           string
	cache             *cache.Cache
	fetchInterval     time.Duration
}
