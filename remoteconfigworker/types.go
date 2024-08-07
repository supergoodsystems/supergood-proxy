package remoteconfigworker

import (
	"time"

	"github.com/supergoodsystems/supergood-proxy/cache"
)

type RemoteConfigOpts struct {
	AdminClientKey string        `yaml:"adminClientKey"`
	BaseURL        string        `yaml:"baseURL"`
	FetchInterval  time.Duration `yaml:"fetchInterval"`
}

type Credential struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TenantConfig struct {
	ClientID     string                  `json:"clientID"`
	ClientSecret string                  `json:"clientSecret"`
	Vendors      map[string][]Credential `json:"vendorConfig"` // map of domain to vendor config
}

type RemoteConfigWorker struct {
	adminClientKey string
	baseURL        string
	cache          *cache.Cache
	fetchInterval  time.Duration
}
