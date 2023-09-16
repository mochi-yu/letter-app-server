package main

import (
	"log"
	"net/http"

	"github.com/mochi-yu/dena-autumn-server/config"
	"github.com/mochi-yu/dena-autumn-server/handler"
	"github.com/rs/cors"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("main: env config load failed: %v", err)
	}
	log.Println(cfg.Port)

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.JsontestHandle)

	// CORSの設定
	c := cors.Default()
	handler := c.Handler(mux)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
