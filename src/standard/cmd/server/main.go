package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	gamev1 "github.com/shwatanap/connect-tic-tac-toe/src/standard/api/game/v1"
	matchingv1 "github.com/shwatanap/connect-tic-tac-toe/src/standard/api/matching/v1"
	handler "github.com/shwatanap/connect-tic-tac-toe/src/standard/handler/server"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	matchingv1.RegisterMatchingServiceServer(server, handler.NewMatchingHandler())
	gamev1.RegisterGameServiceServer(server, handler.NewGameHandler())

	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
