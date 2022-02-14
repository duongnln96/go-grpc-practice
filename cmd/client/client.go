package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/duongnln96/go-grpc-practice/pb/pcbook"
	"github.com/duongnln96/go-grpc-practice/sample"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Start(c *cli.Context) error {
	serverAddress := c.String("address")
	serverPort := c.String("port")
	log.Printf("dial server %s:%s", serverAddress, serverPort)
	serverInfo := fmt.Sprintf("%s:%s", serverAddress, serverPort)

	conn, err := grpc.Dial(serverInfo, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pcbook.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	req := &pcbook.CreateLaptopRequest{
		Laptop: laptop,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			// not a big deal
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return err
	}

	log.Printf("created laptop with id: %s", res.Id)

	return nil
}
