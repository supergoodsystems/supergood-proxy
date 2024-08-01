package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/supergoodsystems/supergood-proxy/cache"
	"github.com/supergoodsystems/supergood-proxy/config"
	"github.com/supergoodsystems/supergood-proxy/proxy"
	"github.com/supergoodsystems/supergood-proxy/remoteconfigworker"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg := config.GetConfigFromEnvironment()
	projectCache := cache.New()

	rcw := remoteconfigworker.New(cfg.RemoteWorkerConfig, &projectCache)
	rp := proxy.New(proxy.ProxyOpts{
		Port:    cfg.ProxyConfig.Port,
		Handler: proxy.NewProxyHandler(&projectCache),
	})

	err := rcw.Start(ctx)
	if err != nil {
		log.Fatalf("Failed to start remote config worker with error: %v", err)
	}
	rp.Start(ctx)

	<-ctx.Done()
	log.Println("Shutting down server...")

	// TODO: Add wait groups to account for both server and worker
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	rp.Stop(shutdownCtx)
	log.Println("Server exiting")
}
