package proxy

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// Proxy is a struct which holds reference to the reverse proxy server
type Proxy struct {
	server *http.Server
}

// ProxyOpts are options to pass to the Proxy constructor New()
type ProxyOpts struct {
	Port string
	Handler *ProxyHandler
}

// New returns a new Reverse Proxy with handler as input
func New(opts ProxyOpts) Proxy {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", opts.Port),
		Handler: http.HandlerFunc(opts.Handler.ServeHTTP),
	}

	return Proxy {
		server: server,
	}
}

// Start begins the reverse proxy
func (p Proxy) Start(ctx context.Context) {
	go func() {
		log.Println("Starting proxy server on :8080")
		if err := p.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
}

// Stop gracefully stops the rever proxy
func (p Proxy) Stop(ctx context.Context){
	if err := p.server.Shutdown(ctx); err != nil {
		log.Fatalf("Proxy shutdown failed with err: %v", err)
	}
}