package auth

import (
	"github.com/duongnln96/go-grpc-practice/pb"
	userRepo "github.com/duongnln96/go-grpc-practice/repo/memory/user"
	jwtSvc "github.com/duongnln96/go-grpc-practice/service/jwt"
)

type ServiceDeps struct {
	UserRepoInstance      userRepo.UserStore
	JwtManagerSvcInstance *jwtSvc.JWTManager
}

type service struct {
	userRepo   userRepo.UserStore
	jwtManager *jwtSvc.JWTManager

	pb.UnimplementedAuthServiceServer
}

// NewLaptopServer returns a new LaptopServer
func NewService(deps ServiceDeps) pb.AuthServiceServer {
	return &service{
		userRepo:   deps.UserRepoInstance,
		jwtManager: deps.JwtManagerSvcInstance,
	}
}
