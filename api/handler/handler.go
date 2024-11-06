package handler

import (
	"context"
	"fmt"

	"github.com/tiksup/tiksup-extractor-worker/internal/database"
	pb "github.com/tiksup/tiksup-extractor-worker/internal/proto"
	"github.com/tiksup/tiksup-extractor-worker/pkg/batch"
)

type Server struct {
	pb.UnimplementedEventTriggerServiceServer
}

func (s *Server) TriggerEvent(ctx context.Context, req *pb.EventRequest) (*pb.EventResponse, error) {
	rdb := batch.RedisRepository{Database: database.RedisClient, CTX: ctx}

	if req.EventName != "next" {
		response := &pb.EventResponse{
			Received: false,
		}
		return response, nil
	}
	response := &pb.EventResponse{
		Received: true,
	}

	fmt.Println(req.String())
	err := rdb.GetMessageQueue(req.UserId)
	if err != nil {
		return nil, err
	}

	return response, nil
}
