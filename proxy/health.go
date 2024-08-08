package proxy

import "net/http"

func NewHealthCheckHandler() *http.ServeMux {
	healthCheckHandler := http.NewServeMux()
	healthCheckHandler.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	return healthCheckHandler
}
