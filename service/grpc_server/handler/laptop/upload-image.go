package laptop

import (
	"bytes"
	"io"
	"log"

	"github.com/duongnln96/go-grpc-practice/pb"
	"github.com/duongnln96/go-grpc-practice/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxImageSize = 1 << 20 // 10

func (s *service) UploadImage(stream pb.LaptopService_UploadImageServer) error {

	req, err := stream.Recv()
	if err != nil {
		return utils.LogError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}

	laptopID := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()
	log.Printf("receive an upload-image request for laptop %s with image type %s", laptopID, imageType)

	laptop, err := s.laptopRepo.Find(laptopID)
	if err != nil {
		return utils.LogError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
	}
	if laptop == nil {
		return utils.LogError(status.Errorf(codes.InvalidArgument, "laptop id %s doesn't exist", laptopID))
	}

	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		if err := utils.ContextError(stream.Context()); err != nil {
			return err
		}

		log.Printf("Waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("no more data")
			break
		}
		if err != nil {
			return utils.LogError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		log.Printf("received a chunk with size: %d", size)

		imageSize += size
		if imageSize > maxImageSize {
			return utils.LogError(status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize))
		}

		// time.Sleep(1 * time.Second) // simulate write date slow
		_, err = imageData.Write(chunk)
		if err != nil {
			return utils.LogError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	imageID, err := s.imageRepo.Save(laptopID, imageType, imageData)
	if err != nil {
		return utils.LogError(status.Errorf(codes.Internal, "cannot save image to the store: %v", err))
	}

	res := &pb.UploadImageResponse{
		Id:   imageID,
		Size: uint32(imageSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return utils.LogError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}

	log.Printf("saved image with id: %s, size: %d", imageID, imageSize)
	return nil
}
