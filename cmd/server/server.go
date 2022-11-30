package server

import (
	"fmt"
	"log"
	"net"

	"github.com/duongnln96/go-grpc-practice/pb"

	imageRepo "github.com/duongnln96/go-grpc-practice/repo/memory/image"
	laptopRepo "github.com/duongnln96/go-grpc-practice/repo/memory/laptop"

	laptopSvc "github.com/duongnln96/go-grpc-practice/service/handler/laptop"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func Start(c *cli.Context) error {
	port := c.Int("port")
	log.Printf("[STARTING] gRPC Server on port %d", port)

	laptopRepoInstance := laptopRepo.NewInMemoryLaptopStore()
	imageRepoInstance := imageRepo.NewDiskImageStore("./tmp/images")

	laptopServer := laptopSvc.NewLaptopServer(laptopSvc.ServiceDeps{
		LaptopRepoInstance: laptopRepoInstance,
		ImageRepoInstance:  imageRepoInstance,
	})

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("[FAIL] Cannot start server: ", err)
	}

	log.Printf("[STARTED] gRPC Server %s", address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("[FAIL] Cannot start server: ", err)
	}

	return nil
}
