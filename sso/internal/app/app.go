package app

import (
	"log/slog"
	"sso/internal/app/grpc"
)

type App struct {
	log *slog.Logger
	gRPCServer *grpc.Server
}