package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	authsvc "sso/internal/services/auth"
	"sso/storage/sqlite"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grcpPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic("failed to init storage: " + err.Error())
	}

	authService := authsvc.New(
		log,
		storage,
		storage,
		storage,
		tokenTTL,
	)

	grpcApp := grpcapp.New(log, authService, grcpPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
