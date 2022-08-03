package main

import (
	"net/http"
	"time"

	"github.com/emurray647/run_log/routing"
)

func main() {
	handler, _ := routing.BuildRoutes()

	srv := &http.Server{
		Handler:      handler,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
