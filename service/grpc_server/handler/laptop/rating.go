package laptop

import (
	"io"
	"log"

	"github.com/duongnln96/go-grpc-practice/pb"
	"github.com/duongnln96/go-grpc-practice/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) RateLaptop(stream pb.LaptopService_RateLaptopServer) error {

	for {
		err := utils.ContextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("no more date")
			break
		}
		if err != nil {
			return utils.LogError(status.Errorf(codes.Internal, "stream.Recv request: %v", err))
		}

		laptopID := req.GetLaptopId()
		score := req.GetScore()

		laptop, err := s.laptopRepo.Find(laptopID)
		if err != nil {
			return utils.LogError(status.Errorf(codes.Internal, "laptopRepo.Find: %v", err))
		}
		if laptop == nil {
			return utils.LogError(status.Errorf(codes.NotFound, "not found laptop_id: %d", laptopID))
		}

		rating, err := s.ratingRepo.Add(laptopID, score)
		if err != nil {
			return utils.LogError(status.Errorf(codes.Internal, "ratingRepo.Add: %v", err))
		}

		response := &pb.RateLaptopResponse{
			LaptopId:     laptopID,
			RatedCount:   rating.Count,
			AverageScore: rating.Sum / float64(rating.Count),
		}
		err = stream.Send(response)
		if err != nil {
			return utils.LogError(status.Errorf(codes.Unknown, "stream.Send %v", err))
		}
	}

	return nil
}
