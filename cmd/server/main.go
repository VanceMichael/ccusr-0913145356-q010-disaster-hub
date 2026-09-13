package main

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok", "service": "q010-mudslide-hub"})
	})
	mux.Handle("/metrics", promhttp.Handler())
	return mux
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("服务监听端口 %s", port)
	log.Fatal(http.ListenAndServe(":"+port, newHandler()))
}
