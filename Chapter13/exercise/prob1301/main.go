package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := createChiRouter()
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

func createServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		t := time.Now().Format(time.RFC3339)
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(t))
		return
	})
	return mux
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
