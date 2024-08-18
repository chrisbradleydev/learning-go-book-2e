package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := createChiRouter()
	// r := createServeMux()
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func createChiRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().Format(time.RFC3339)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(t))
	})
	return r
}

func createServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		t := time.Now().Format(time.RFC3339)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(t))
		return
	})
	return mux
}
