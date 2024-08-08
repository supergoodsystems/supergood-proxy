package proxy

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// Proxy is a struct which holds reference to the reverse proxy server
type Proxy struct {
	server            *http.Server
	healthCheckServer *http.Server
}

// ProxyOpts are options to pass to the Proxy constructor New()
type ProxyOpts struct {
	Port               string
	HealthCheckPort    string
	Handler            *ProxyHandler
	HealthCheckHandler *http.ServeMux
}

// New returns a new Reverse Proxy with handler as input
func New(opts ProxyOpts) Proxy {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", opts.Port),
		Handler: http.HandlerFunc(opts.Handler.ServeHTTP),
	}

	healthCheck := &http.Server{
		Addr:    fmt.Sprintf(":%s", opts.HealthCheckPort),
		Handler: opts.HealthCheckHandler,
	}

	return Proxy{
		server:            server,
		healthCheckServer: healthCheck,
	}
}

// Start begins the reverse proxy
func (p Proxy) Start(ctx context.Context) {
	go func() {
		log.Println("Starting proxy server")
		if err := p.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting proxy server: %v", err)
		}
	}()
	go func() {
		log.Println("Starting health check server")
		if err := p.healthCheckServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting health check server: %v", err)
		}
	}()
}

// Stop gracefully stops the rever proxy
func (p Proxy) Stop(ctx context.Context) {
	if err := p.server.Shutdown(ctx); err != nil {
		log.Fatalf("Proxy shutdown failed with err: %v", err)
	}
	if err := p.healthCheckServer.Shutdown(ctx); err != nil {
		log.Fatalf("Health check server shutdown failed with err: %v", err)
	}
}
