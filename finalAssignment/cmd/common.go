package cmd

import (
	_ "embed"
	"net/http"

	"github.com/flowchartsman/swaggerui"
)

//go:embed swagger.json
var spec []byte

func CreateCommonMux(h http.Handler) http.Handler {
	r := http.NewServeMux()
	fs := http.FileServer(http.Dir("./client/build/static"))
	r.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(spec)))
	r.HandleFunc("/app", index)
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	r.Handle("/api", h)

	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/build/index.html")
}
