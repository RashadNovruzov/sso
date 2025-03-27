package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"syscall"

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

	application := app.New(log, config.GRPC.Port, config.StoragePath, config.TokenTTL)
	go application.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop // waiting for one of those signals which written above from stop channel. After receiving signal code will continue and shutdown gracefully

	application.GRPCServ.Stop()

	log.Info("Application stopped")

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
