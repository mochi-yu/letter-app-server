package main

import (
	"context"
	"log"
	"net/http"

	"github.com/mochi-yu/dena-autumn-server/config"
	"github.com/rs/cors"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("main: env config load failed: %v", err)
	}

	mux, cleanup, err := NewMux(context.Background(), cfg)
	defer cleanup()
	if err != nil {
		log.Fatalf("main: NewMux error: %v", err)
	}

	// CORSの設定
	c := cors.Default()
	corsHandler := c.Handler(mux)

	log.Println("listen: ", cfg.Port)
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
