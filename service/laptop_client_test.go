package service_test

import (
	"context"
	"net"
	"testing"

	"github.com/duongnln96/go-grpc-practice/pb/pcbook"
	"github.com/duongnln96/go-grpc-practice/sample"
	"github.com/duongnln96/go-grpc-practice/serializer"
	"github.com/duongnln96/go-grpc-practice/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopStore := service.NewInMemoryLaptopStore()
	serverAddress := startTestLaptopServer(t, laptopStore)
	client := newTestLaptopClient(t, serverAddress)

	laptopReq := sample.NewLaptop()
	expectedID := laptopReq.Id
	req := &pcbook.CreateLaptopRequest{
		Laptop: laptopReq,
	}

	res, err := client.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	laptopInDb, err := laptopStore.Find(res.Id)
	require.NoError(t, err)

	requireSameLaptop(t, laptopReq, laptopInDb)
}

func startTestLaptopServer(t *testing.T, laptopStore service.LaptopStore) string {
	laptopServer := service.NewLaptopServer(laptopStore)

	grpcServer := grpc.NewServer()
	pcbook.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pcbook.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)

	return pcbook.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *pcbook.Laptop, laptop2 *pcbook.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
