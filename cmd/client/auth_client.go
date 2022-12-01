package client

import (
	"context"
	"time"

	"github.com/duongnln96/go-grpc-practice/pb"
	"google.golang.org/grpc"
)

type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}

func (client *AuthClient) Login() (accessToken string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, errRet := client.service.Login(ctx, req)
	if errRet != nil {
		err = errRet
		return
	}

	accessToken = res.GetAccessToken()

	return
}
