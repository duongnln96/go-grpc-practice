package laptop

import (
	"github.com/duongnln96/go-grpc-practice/pb"
	imageRepo "github.com/duongnln96/go-grpc-practice/repo/memory/image"
	latopRepo "github.com/duongnln96/go-grpc-practice/repo/memory/laptop"
)

type ServiceDeps struct {
	LaptopRepoInstance latopRepo.LaptopStore
	ImageRepoInstance  imageRepo.ImageStore
}

type service struct {
	laptopRepo latopRepo.LaptopStore
	imageRepo  imageRepo.ImageStore

	pb.UnimplementedLaptopServiceServer
}

// NewLaptopServer returns a new LaptopServer
func NewService(deps ServiceDeps) pb.LaptopServiceServer {
	return &service{
		laptopRepo: deps.LaptopRepoInstance,
		imageRepo:  deps.ImageRepoInstance,
	}
}
