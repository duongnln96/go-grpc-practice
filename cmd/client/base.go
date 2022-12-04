package client

import (
	"fmt"
	"log"
	"time"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

const (
	username        = "admin1"
	password        = "secret"
	refreshDuration = 30 * time.Second
)

func authMethods() map[string]bool {
	const laptopServicePath = "/duongnln.pcbook.LaptopService/"

	return map[string]bool{
		fmt.Sprintf("%s%s", laptopServicePath, "CreateLaptop"): true,
		fmt.Sprintf("%s%s", laptopServicePath, "UploadImage"):  true,
		fmt.Sprintf("%s%s", laptopServicePath, "RateLaptop"):   true,
	}
}

func Start(c *cli.Context) error {
	serverAddress := c.String("address")
	serverPort := c.String("port")

	serverInfo := fmt.Sprintf("%s:%s", serverAddress, serverPort)
	log.Printf("Dialing gRPC server %s", serverInfo)

	// create connection to server
	cc1, err := grpc.Dial(serverInfo, grpc.WithInsecure())
	if err != nil {
		log.Fatal("[FAIL] Cannot dial server: ", err)
	}

	authClient := NewAuthClient(cc1, username, password)

	interceptor, err := NewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	cc2, err := grpc.Dial(
		serverInfo,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := NewLaptopClient(cc2)

	testRatingLaptop(laptopClient)

	return nil
}
