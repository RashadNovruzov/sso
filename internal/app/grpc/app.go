package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
	authgrpc "sso/internal/grpc/auth"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int, authservice authgrpc.Auth) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer, authservice)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.run(); err != nil {
		panic("Grpc Server not started " + err.Error())
	}
}

func (a *App) run() error {
	const op = "grpcapp.Run"

	log := a.log.With(slog.String("op", op))

	log.Info("Starting GRPC server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("GRPC server started ", slog.String("address", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	log := a.log.With(slog.String("op", op))

	log.Info("Stopping GRPC server")

	a.gRPCServer.GracefulStop()

}
