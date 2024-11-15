package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/exoneges/doodocs-days-backend/internal/config"
	"github.com/exoneges/doodocs-days-backend/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	handler.SetMuxArchiveHanle(mux)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%v", config.PORT),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	slog.Info("Starting srver on", "Port", config.PORT, "Dir", config.DIR, "ProtectedIP", config.PROTECTED)
	defer config.LOGFILE.Close()

	if err := server.ListenAndServe(); err != nil {
		slog.Error("Error starting server:", "err", err)
	}
}
