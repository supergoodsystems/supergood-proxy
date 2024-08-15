package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/supergoodsystems/supergood-proxy/cache"
)

const SupergoodUpstreamHeader = "X-Supergood-Upstream"
const SupergoodClientIDHeader = "X-Supergood-ClientID"
const SupergoodClientSecretHeader = "X-Supergood-ClientSecret"

// ProxyHandler is an HTTP handler that proxies requests to another server
type ProxyHandler struct {
	projectCache *cache.Cache
}

// NewProxyHandler creates a new ProxyHandler with a projectCache as required input
func NewProxyHandler(projectCache *cache.Cache) *ProxyHandler {
	return &ProxyHandler{
		projectCache: projectCache,
	}
}

// ServeHTTP is the proxy handler which will stuff credentials into the
// proxied request based off clientID, clientSecret, fqdn stored in request headers
func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upstream := r.Header.Get(SupergoodUpstreamHeader)
	targetURL, err := url.Parse(upstream)

	if err != nil {
		http.Error(w, fmt.Sprintf("Supergood: Unable to parse upstream URL:%s", upstream), http.StatusBadRequest)
		return
	}

	clientID := r.Header.Get(SupergoodClientIDHeader)
	clientSecret := r.Header.Get(SupergoodClientSecretHeader)

	projectConfig := p.projectCache.Get(clientID)
	if projectConfig == nil || projectConfig.ClientSecret != clientSecret {
		http.Error(w, "Invalid Supergood Credentials", http.StatusUnauthorized)
		return
	}

	director := func(req *http.Request) {
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.Host = targetURL.Host

		req.Header.Del(SupergoodClientIDHeader)
		req.Header.Del(SupergoodClientSecretHeader)

		vendorConfig, ok := projectConfig.Vendors[targetURL.Host]
		if !ok {
			return
		}
		for _, cred := range vendorConfig.Credentials {
			req.Header.Del(cred.Key)
			req.Header.Add(cred.Key, cred.Value)
		}
	}
	proxy := &httputil.ReverseProxy{Director: director}

	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		log.Printf("Error during proxying: %v", err)
		http.Error(rw, "Supergood proxy error", http.StatusBadGateway)
	}

	proxy.ServeHTTP(w, r)
}
