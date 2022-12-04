package laptop

import (
	"github.com/duongnln96/go-grpc-practice/pb"
	imageRepo "github.com/duongnln96/go-grpc-practice/repo/memory/image"
	latopRepo "github.com/duongnln96/go-grpc-practice/repo/memory/laptop"
	ratingRepo "github.com/duongnln96/go-grpc-practice/repo/memory/rating"
)

type ServiceDeps struct {
	LaptopRepoInstance latopRepo.LaptopStore
	ImageRepoInstance  imageRepo.ImageStore
	RatingRepoInstance ratingRepo.RatingStore
}

type service struct {
	laptopRepo latopRepo.LaptopStore
	imageRepo  imageRepo.ImageStore
	ratingRepo ratingRepo.RatingStore

	pb.UnimplementedLaptopServiceServer
}

// NewLaptopServer returns a new LaptopServer
func NewService(deps ServiceDeps) pb.LaptopServiceServer {
	return &service{
		laptopRepo: deps.LaptopRepoInstance,
		imageRepo:  deps.ImageRepoInstance,
		ratingRepo: deps.RatingRepoInstance,
	}
}
