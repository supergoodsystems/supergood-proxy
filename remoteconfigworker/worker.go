package remoteconfigworker

import (
	"context"
	"log"
	"time"

	"github.com/supergoodsystems/supergood-proxy/cache"
)

// New creates a new remote config worker
func New(opts RemoteConfigOpts, tenantCache *cache.Cache) RemoteConfigWorker {
	return RemoteConfigWorker{
		baseURL:           opts.BaseURL,
		adminClientId:     opts.AdminClientID,
		adminClientSecret: opts.AdminClientSecret,
		cache:             tenantCache,
		fetchInterval:     opts.FetchInterval,
	}
}

// Start fetches the the remote config cache and begins the worker
func (rc *RemoteConfigWorker) Start(ctx context.Context) error {
	err := rc.fetchAndSetConfig()
	if err != nil {
		return err
	}
	go rc.Refresh(ctx)
	return nil
}

// RefreshRemoteConfig refreshes the remote config on an interval
// and receives a close channel to gracefully return on application exit
func (rc *RemoteConfigWorker) Refresh(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Gracefully exiting Remote Config Worker")
			return
		case <-time.After(rc.fetchInterval):
			if err := rc.fetchAndSetConfig(); err != nil {
				log.Println("Failed to fetch config")
			}
		}
	}
}

// fetchAndSetConfig fetches the remote config from the supergood /proxy-config endpoint
// and then sets it in the Cache on the RemoteConfig
func (rc *RemoteConfigWorker) fetchAndSetConfig() error {
	resp, err := rc.fetch()
	if err != nil {
		return err
	}

	for _, config := range resp {
		cacheVal := responseToCacheVal(config)
		rc.cache.Set(config.ClientID, &cacheVal)
	}
	return nil
}

// responseToCacheVal marshals the TenantConfig response object into a cache value.
// I'd like to not to have to convert TenantConfig into cache.CacheVal, but remoteconfigworker
// and cache are separate packages and I dont want the cache package to have a dependency
// on the remoteconfigworker package. There's probably a better way, most likely moving these
// struct definitions to a shared package - but not sure thats good go.
func responseToCacheVal(config TenantConfig) cache.CacheVal {
	cacheVal := cache.CacheVal {
		ClientID: config.ClientID,
		ClientSecret: config.ClientSecret,
		Vendors: map[string]cache.VendorConfig{}, 
	}

	for domain, config := range config.Vendors {
		cacheCreds := []cache.Credential{}
		for _, cred := range config.Credentials {
			cacheCreds = append(cacheCreds, cache.Credential{Key: cred.Key, Value: cred.Value})
		}
		cacheVal.Vendors[domain] = cache.VendorConfig{
			Credentials: cacheCreds,
		}
	}
	return cacheVal
}
