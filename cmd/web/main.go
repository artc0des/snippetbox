package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	//creating commnand-line flag that defines a port to be used in env
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	//structured logger definition
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	//JSON logger syntax with debug level enabled and source of log
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))
	jsonLogger.Debug("Debug log output test")

	app := &application{
		logger: logger,
	}

	logger.Info("Starting server", "port", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
