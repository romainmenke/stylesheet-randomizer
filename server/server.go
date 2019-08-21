package server

import (
	mr "math/rand"
	"net/http"
	"strconv"

	"github.com/romainmenke/stylesheet-randomizer/randomstylesheet"
)

func Serve(addr string) *http.Server {
	mux := http.NewServeMux()

	mux.Handle("/stylesheet.css", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "max-age=0, private")
		w.Header().Add("Content-Type", "text/css")

		seed, err := strconv.ParseInt(r.URL.Query().Get("seed"), 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		randomstylesheet.Write(w, seed)
	}))

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "max-age=0, private")
		w.Header().Add("Content-Type", "text/html")

		seed, err := strconv.ParseInt(r.URL.Query().Get("seed"), 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError)+" : invalid or missing \"seed\" query parameter", http.StatusInternalServerError)
			return
		}

		mr.Seed(seed)

		writeSampleHTML(w, seed)
	}))

	server := &http.Server{
		Handler: mux,
		Addr:    addr,
	}

	return server
}
