package cache

import "sync"

type Cache struct {
	cache  map[string]*CacheVal
	mutex             *sync.RWMutex
}

type CacheVal struct {
	ClientID string
	ClientSecret string
	Vendors map[string]VendorConfig
}

type VendorConfig struct {
	Credentials []Credential
}

type Credential struct {
	Key string
	Value string
}