package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/tiksup/tiksup-extractor-worker/api/handler"
	"github.com/tiksup/tiksup-extractor-worker/internal/config"
	"github.com/tiksup/tiksup-extractor-worker/internal/database"
	pb "github.com/tiksup/tiksup-extractor-worker/internal/proto"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("\033[31mNot .env file found. Using system variables\033[0m")
	}

	if err := database.GetRedisConnection(context.Background()); err != nil {
		log.Fatalf("Error trying connect to redis: %v\n", err)
	}

	if err := database.GetMongoConnection(context.TODO()); err != nil {
		log.Fatalf("Error trying connect to mongo: %v\n", err)
	}
}

func main() {
	PORT := os.Getenv("PORT")
	TARGET_GRPC_SERVER_HOST := os.Getenv("TARGET_GRPC_SERVER_HOST")
	TARGET_GRPC_SERVER_PORT := os.Getenv("TARGET_GRPC_SERVER_PORT")
	if PORT == "" {
		log.Fatal("PORT not set in environment variables")
	}
	if TARGET_GRPC_SERVER_HOST == "" || TARGET_GRPC_SERVER_PORT == "" {
		log.Fatal("TARGET_GRPC_SERVER not set in environment variables")
	}

	addr := fmt.Sprintf(":%s", PORT)
	grpcTargetServer := fmt.Sprintf("%s:%s", TARGET_GRPC_SERVER_HOST, TARGET_GRPC_SERVER_PORT)

	defer func() {
		database.RedisClient.Close()
		database.MongoClient.Disconnect(context.TODO())
	}()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	client, err := config.CreateInteractionClient(grpcTargetServer)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer client.Conn.Close()

	s := grpc.NewServer()
	pb.RegisterEventTriggerServiceServer(s, &handler.Server{
		Client: client.Client,
	})

	log.Printf("\033[32mSERVER LISTEN ON PORT %s\033[0m", PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
