package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	options := &slog.HandlerOptions{}
	handler := slog.NewJSONHandler(os.Stderr, options)
	mySlog := slog.New(handler)
	r := createChiRouter(mySlog)
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}
	err := s.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}

func createChiRouter(logger *slog.Logger) chi.Router {
	r := chi.NewRouter().With(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			var ip string
			if strings.Contains(req.RemoteAddr, "[::1]") {
				ip = "127.0.0.1"
			} else {
				ip, _, _ = strings.Cut(req.RemoteAddr, ":")
			}
			logger.Info("incoming IP", "ip", ip)
			handler.ServeHTTP(rw, req)
		})
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().Format(time.RFC3339)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(t))
	})
	return r
}
