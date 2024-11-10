package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/tiksup/tiksup-extractor-worker/internal/database"
	pb "github.com/tiksup/tiksup-extractor-worker/internal/proto"
	"github.com/tiksup/tiksup-extractor-worker/pkg/batch"
	"github.com/tiksup/tiksup-extractor-worker/pkg/movie"
)

type Server struct {
	pb.UnimplementedEventTriggerServiceServer
	Client pb.InteractionsServiceClient
}

func (s *Server) TriggerEvent(ctx context.Context, req *pb.EventRequest) (*pb.EventResponse, error) {
	rdb := &batch.RedisRepository{Database: database.RedisClient, CTX: ctx}
	movieRepository := &movie.MongoRepository{Database: database.MongoDatabase, CTX: ctx}

	if req.EventName != "next" {
		return &pb.EventResponse{Received: false}, nil
	}

	queue, err := rdb.GetMessageQueue(req.UserId)
	if err != nil {
		log.Printf("Error getting message queue: %v\n", err)
		return nil, err
	}

	movies, err := movieRepository.GetMoviesExcludeHistory(req.UserId)
	if err != nil {
		log.Printf("Error getting movies: %v", err)
		return nil, err
	}

	data := batch.PreprocessedData{
		User:   req.UserId,
		Data:   queue,
		Movies: movies,
	}

	err = s.SendInteracctions(ctx, data)
	if err != nil {
		log.Printf("Error getting movies: %v", err)
		return nil, err
	}

	log.Println("Data received")
	return &pb.EventResponse{Received: true}, nil
}

func (s *Server) SendInteracctions(ctx context.Context, data batch.PreprocessedData) error {
	userInfo := make([]*pb.UserInfoRequest, len(data.Data))
	movies := make([]*pb.MoviesRequest, len(data.Movies))

	for i, d := range data.Data {
		userInfo[i] = &pb.UserInfoRequest{
			MovieId:        d.MovieID,
			WatchingTime:   d.WatchingTime,
			WatchingRepeat: d.WatchingRepeat,
			Interactions: &pb.InteractionsRequest{
				Genre:       d.Interactions.Genre,
				Protagonist: d.Interactions.Protagonist,
				Director:    d.Interactions.Director,
			},
		}
	}

	for i, d := range data.Movies {
		movies[i] = &pb.MoviesRequest{
			Id:          d.ID.Hex(),
			Url:         d.URL,
			Title:       d.Title,
			Genre:       d.Genre,
			Protagonist: d.Protagonist,
			Director:    d.Director,
		}
	}

	request := &pb.PreprocessedDataRequest{
		User:   data.User,
		Data:   userInfo,
		Movies: movies,
	}

	res, err := s.Client.ProcessData(ctx, request)
	if err != nil {
		return fmt.Errorf("Error to sending request: %w", err)
	}
	if !res.Success {
		return fmt.Errorf("Wrong data")
	}

	return nil
}
