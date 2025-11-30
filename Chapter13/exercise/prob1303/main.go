package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
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
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func createChiRouter(logger *slog.Logger) chi.Router {
	r := chi.NewRouter().With(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ip, _, _ := strings.Cut(req.RemoteAddr, ":")
			logger.Info("incoming IP", "ip", ip)
			handler.ServeHTTP(rw, req)
		})
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		var out string
		if r.Header.Get("Accept") == "application/json" {
			out = buildJSON(now)
		} else {
			out = buildText(now)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(out))
	})
	return r
}

func buildText(now time.Time) string {
	return now.Format(time.RFC3339)
}

func buildJSON(now time.Time) string {
	timeOut := struct {
		DayOfWeek  string `json:"day_of_week"`
		DayOfMonth int    `json:"day_of_month"`
		Month      string `json:"month"`
		Year       int    `json:"year"`
		Hour       int    `json:"hour"`
		Minute     int    `json:"minute"`
		Second     int    `json:"second"`
	}{
		DayOfWeek:  now.Weekday().String(),
		DayOfMonth: now.Day(),
		Month:      now.Month().String(),
		Year:       now.Year(),
		Hour:       now.Hour(),
		Minute:     now.Minute(),
		Second:     now.Second(),
	}
	out, _ := json.Marshal(timeOut)
	return string(out)
}
