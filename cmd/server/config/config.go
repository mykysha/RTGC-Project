package config

import (
	"net/http"
	"time"
)

func newConfigServer() *http.Server {
	srv := http.Server{
		Addr:              ":8080",
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       3 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	return &srv
}
