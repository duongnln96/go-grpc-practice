package auth

import (
	"context"

	"github.com/duongnln96/go-grpc-practice/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.userRepo.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "userRepo.Find: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.Password) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := s.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.LoginResponse{
		AccessToken: token,
	}

	return res, nil
}
