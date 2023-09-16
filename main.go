package main

import (
	"fmt"
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
	log.Println(cfg.Port)

	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandle)

	// CORSの設定
	c := cors.Default()
	handler := c.Handler(mux)

	log.Fatal(http.ListenAndServe(":8080", handler))

}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintf(w, "{\"hello\": \"world\"}"); err != nil {
		fmt.Printf("helloHandle write error: %v", err)
	}
}
