/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/supergoodsystems/supergood-proxy/cache"
	"github.com/supergoodsystems/supergood-proxy/config"
	"github.com/supergoodsystems/supergood-proxy/proxy"
	"github.com/supergoodsystems/supergood-proxy/remoteconfigworker"
)

func run() error {
	path := ""
	cmd := &cobra.Command{
		Use:           "supergood-proxy",
		Short:         "Run Supergood Proxy",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(_ *cobra.Command, _ []string) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			defer cancel()

			cfg := config.GetConfig(path)
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
			return nil
		},
	}

	cmd.PersistentFlags().StringVar(
		&path,
		"file",
		path,
		"Path to a file where '.yml' configuration is stored",
	)

	return cmd.Execute()
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
