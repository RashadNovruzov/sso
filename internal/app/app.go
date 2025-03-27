package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	GRPCServ *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {

	//Todo: init storage

	//Todo: init auth service

	// Init grpc app

	return &App{
		GRPCServ: grpcapp.New(log, grpcPort),
	}
}

func (a *App) MustRun() {
	a.GRPCServ.MustRun()
}
