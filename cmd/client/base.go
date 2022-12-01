package client

import (
	"fmt"
	"log"
	"time"

	"github.com/duongnln96/go-grpc-practice/pb"
	sample "github.com/duongnln96/go-grpc-practice/sample/generator"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func testCreateLaptop(laptopClient *LaptopClient) {
	laptopClient.CreateLaptop(sample.NewLaptop())
}

func testSearchLaptop(laptopClient *LaptopClient) {
	for i := 0; i < 10; i++ {
		laptopClient.CreateLaptop(sample.NewLaptop())
	}

	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}

	laptopClient.SearchLaptop(filter)
}

func testUploadImage(laptopClient *LaptopClient) {
	laptop := sample.NewLaptop()
	laptopClient.CreateLaptop(laptop)
	laptopClient.UploadImage(laptop.GetId(), "./sample/images/9080679.jpeg")
}

const (
	username        = "admin1"
	password        = "secret"
	refreshDuration = 30 * time.Second
)

func authMethods() map[string]bool {
	const laptopServicePath = "/duongnln.pcbook.LaptopService/"

	return map[string]bool{
		laptopServicePath + "CreateLaptop": true,
		laptopServicePath + "UploadImage":  true,
		laptopServicePath + "RateLaptop":   true,
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

	testCreateLaptop(laptopClient)

	return nil
}
