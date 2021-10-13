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
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       readTime,
		ReadHeaderTimeout: readerHeaderTime,
		WriteTimeout:      writeTime,
		IdleTimeout:       idleTime,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	return &srv
}
