package main

import (
	"context"

	"github.com/felipeversiane/go-otel/internal/infra/config"
	"github.com/felipeversiane/go-otel/internal/infra/config/log"
	"github.com/felipeversiane/go-otel/internal/infra/server"
	"github.com/felipeversiane/go-otel/internal/infra/services/database"
	"github.com/felipeversiane/go-otel/internal/infra/services/observability"
)

func main() {
	cfg := config.NewConfig()

	logger := log.NewLogger(cfg.GetLogConfig())
	logger.Configure()

	ctx := context.Background()
	otel := observability.NewOtel(observability.OtelConfig{
		ServiceName:              "go-otel",
		ServiceVersion:           "0.0.1",
		OtelExporterOtlpEndpoint: "otel-collector:4317",
		OtelExporterOtlpInsecure: true,
	})
	otel.SetupOtel(ctx)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database := database.NewDatabaseConnection(ctx, cfg.GetDatabaseConfig())
	defer database.Close()

	server := server.NewServer(cfg.GetServerConfig(), database)
	server.SetupRouter()
	server.Start()
}
