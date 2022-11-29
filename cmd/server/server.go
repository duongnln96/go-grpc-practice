package server

import (
	"fmt"
	"log"
	"net"

	"github.com/duongnln96/go-grpc-practice/pb"
	"github.com/duongnln96/go-grpc-practice/service"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func Start(c *cli.Context) error {
	port := c.Int("port")
	log.Printf("start server on port %d", port)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("images")
	laptopServer := service.NewLaptopServer(laptopStore, imageStore)

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	return nil
}
