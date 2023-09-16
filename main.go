package main

import (
	"context"
	"log"
	"net/http"

	"github.com/mochi-yu/dena-autumn-server/config"
	"github.com/rs/cors"
)

func main() {

	// str := `[{"points": [{"x": 0.5887799693822837, "y": 0.35986568660338486, "time": "6171385531390910809", "pressure": 1}, {"x": 0.6926764359256931, "y": 0.13428400747097005, "time": "1230306591142793169", "pressure": 1}], "dotSize": 0.2841155558975101, "maxWidth": 0.7105549050363199, "minWidth": 0.4753475039759988, "penColor": "black", "compositeOperation": "source-over", "velocityFilterWeight": -0.5775088726519758}]`
	// byt := []byte(str)
	// var t entity.Strokes
	// if err := json.Unmarshal(byt, &t); err != nil {
	// 	log.Fatalf("muri")
	// }
	// fmt.Println(t)

	// bb, _ := json.Marshal(t)
	// fmt.Println(string(bb))

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
