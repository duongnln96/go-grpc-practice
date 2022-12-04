package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/duongnln96/go-grpc-practice/pb"

	imageRepo "github.com/duongnln96/go-grpc-practice/repo/memory/image"
	laptopRepo "github.com/duongnln96/go-grpc-practice/repo/memory/laptop"
	ratingRepo "github.com/duongnln96/go-grpc-practice/repo/memory/rating"
	userRepo "github.com/duongnln96/go-grpc-practice/repo/memory/user"

	authSvc "github.com/duongnln96/go-grpc-practice/service/grpc_server/handler/auth"
	laptopSvc "github.com/duongnln96/go-grpc-practice/service/grpc_server/handler/laptop"
	"github.com/duongnln96/go-grpc-practice/service/grpc_server/interceptor"

	jwtSvc "github.com/duongnln96/go-grpc-practice/service/jwt"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

func accessibleRoles() map[string][]string {
	const laptopServicePath string = "/duongnln.pcbook.LaptopService/"
	return map[string][]string{
		fmt.Sprintf("%s%s", laptopServicePath, "CreateLaptop"): {"admin"},
		fmt.Sprintf("%s%s", laptopServicePath, "UploadImage"):  {"admin"},
		fmt.Sprintf("%s%s", laptopServicePath, "RateLaptop"):   {"admin", "user"},
	}
}

func Start(c *cli.Context) error {
	port := c.Int("port")
	log.Printf("[STARTING] gRPC Server on port %d", port)

	// init repo
	laptopRepoInstance := laptopRepo.NewInMemoryLaptopStore()

	ratingRepoInstance := ratingRepo.NewInMemoryRatingStore()

	imageRepoInstance := imageRepo.NewDiskImageStore("./tmp/")

	userRepoInstance := userRepo.NewInMemoryUserStore()
	err := userRepoInstance.SeedUser()
	if err != nil {
		log.Fatal("[FAIL] Cannot seed users: ", err)
	}

	// init services
	jwtService := jwtSvc.NewJWTManager(secretKey, tokenDuration)

	laptopService := laptopSvc.NewService(laptopSvc.ServiceDeps{
		LaptopRepoInstance: laptopRepoInstance,
		ImageRepoInstance:  imageRepoInstance,
		RatingRepoInstance: ratingRepoInstance,
	})

	authService := authSvc.NewService(authSvc.ServiceDeps{
		UserRepoInstance:      userRepoInstance,
		JwtManagerSvcInstance: jwtService,
	})

	interceptorInstance := interceptor.NewAuthInterceptor(jwtService, accessibleRoles())

	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptorInstance.Unary()),
		grpc.StreamInterceptor(interceptorInstance.Stream()),
	}
	grpcServer := grpc.NewServer(serverOptions...)
	pb.RegisterLaptopServiceServer(grpcServer, laptopService)
	pb.RegisterAuthServiceServer(grpcServer, authService)

	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("[FAIL] Cannot start server: ", err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-sigCh
		log.Printf("Got signal %v, attempting graceful shutdown", s)
		grpcServer.GracefulStop()
		// grpc.Stop() // leads to error while receiving stream response: rpc error: code = Unavailable desc = transport is closing
		wg.Done()
	}()

	log.Printf("[STARTED] gRPC Server %s", address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("[FAIL] Cannot start server: ", err)
	}

	wg.Wait()
	log.Println("[SHUTDOWN] Clean shutdown")
	return nil
}
