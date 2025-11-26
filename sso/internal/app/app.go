package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grcpPort int,
	storagePath string,
	tolenTTL time.Duration,
) *App {
	grpcApp := grpcapp.New(log, grcpPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}