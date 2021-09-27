package config

import (
	"net/http"
	"time"
)

func newConfigServer() *http.Server {
	var (
		readTime         = 1 * time.Second
		writeTime        = 5 * time.Second
		idleTime         = 3 * time.Second
		readerHeaderTime = 2 * time.Second
	)

	srv := http.Server{
		Addr:              ":8080",
		ReadTimeout:       readTime,
		WriteTimeout:      writeTime,
		IdleTimeout:       idleTime,
		ReadHeaderTimeout: readerHeaderTime,
	}

	return &srv
}
