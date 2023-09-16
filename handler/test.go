package handler

import "net/http"

func JsonTest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	test := struct {
		Hello string `json:"Hello"`
	}{
		Hello: "World",
	}

	RespondJSON(ctx, w, test, http.StatusOK)
}
