package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"github.com/matthewjamesboyle/logging-module/internal/db"
	"github.com/matthewjamesboyle/logging-module/internal/elasticsearch"
	"github.com/matthewjamesboyle/logging-module/internal/library"
	ilog "github.com/matthewjamesboyle/logging-module/internal/log"
	"github.com/matthewjamesboyle/logging-module/internal/transport"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	logLevel := slog.LevelError

	esURL := os.Getenv("ENV_ES_URL")
	if esURL == "" {
		log.Fatal("elasticsearch url cannot be empty")
	}
	logMode := os.Getenv("ENV_LOG_LEVEL")
	if logMode == "debug" {
		logLevel = slog.LevelInfo
	}

	ctx := context.Background()

	d := db.MockDb{}

	es, err := elasticsearch.NewESWriter(esURL)
	if err != nil {
		log.Fatal(err)
	}

	loggers := []io.Writer{es}

	l := ilog.NewMultiSourceLoggerLogger(&slog.HandlerOptions{
		Level: logLevel,
	}, loggers...)

	a := library.NewMockAdaptor(d)

	sa := map[string]struct{}{
		"james smith":   {},
		"jack jones":    {},
		"rachel barnes": {},
	}

	svc, err := library.NewService(a, sa, l)
	if err != nil {
		l.ErrorContext(ctx, "failed to create new service", slog.Any("err", err))
		os.Exit(1)
	}

	h, err := transport.NewHandler(*svc, l)
	if err != nil {
		l.ErrorContext(ctx, "failed to create new handler", slog.Any("err", err))
		os.Exit(1)
	}

	mux := transport.NewMux(*h)

	l.InfoContext(ctx, "server started")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		l.ErrorContext(ctx, "server stopped", slog.Any("err", err))
		os.Exit(1)
	}
}
