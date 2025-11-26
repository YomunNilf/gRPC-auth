package auth

import (
	"context"

	ssov1 "github.com/YomunNilf/gRPC-auth/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type ServerAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &ServerAPI{})
}

func (s *ServerAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginRequest, error) {
	panic("implement me")
}

func (s *ServerAPI) IsAdmin(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginRequest, error) {
	panic("implement me")
}

func (s *ServerAPI) Register(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginRequest, error) {
	panic("implement me")
}