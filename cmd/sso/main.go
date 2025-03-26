package main

import (
	"log/slog"
	"os"
	"sso/internal/config"

	"github.com/RashadNovruzov/prettyslogger/handlers/slogpretty"
)

const (
	envlocal = "local"
	envProd  = "prod"
)

func main() {
	// Todo: init object config
	config := config.MustLoad()

	log := setupLogger(config.Env)

	log.Info("starting application",
		slog.String("env", config.Env),
		slog.Any("config", config),
	)

	log.Debug("Debug log")

	log.Error("Error log")

	log.Warn("Warning log")

	// Todo: initialize (app)

	// Todo: run grpc server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envlocal:
		log = setupPrettyLogger()
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettyLogger() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
