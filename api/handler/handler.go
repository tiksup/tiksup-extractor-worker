package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tiksup/tiksup-extractor-worker/internal/database"
	pb "github.com/tiksup/tiksup-extractor-worker/internal/proto"
	"github.com/tiksup/tiksup-extractor-worker/pkg/batch"
	"github.com/tiksup/tiksup-extractor-worker/pkg/movie"
)

type Server struct {
	pb.UnimplementedEventTriggerServiceServer
}

func (s *Server) TriggerEvent(ctx context.Context, req *pb.EventRequest) (*pb.EventResponse, error) {
	rdb := &batch.RedisRepository{Database: database.RedisClient, CTX: ctx}
	movieRepository := &movie.MongoRepository{Database: database.MongoDatabase, CTX: ctx}

	if req.EventName != "next" {
		return &pb.EventResponse{Received: false}, nil
	}

	data, err := rdb.GetMessageQueue(req.UserId)
	if err != nil {
		log.Printf("Error getting message queue: %v\n", err)
		return nil, err
	}

	movies, err := movieRepository.GetMoviesExcludeHistory(req.UserId)
	if err != nil {
		log.Printf("Error getting movies: %v", err)
		return nil, err
	}

	data2 := batch.PreprocessedData{
		User:   req.UserId,
		Data:   data,
		Movies: movies,
	}

	d, _ := json.MarshalIndent(data2, "", "    ")
	fmt.Printf("%s\n", d)

	log.Println("Data received")

	return &pb.EventResponse{Received: true}, nil
}
