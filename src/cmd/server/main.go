package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	gamev1connect "github.com/shwatanap/connect-tic-tac-toe/src/api/game/v1/gamev1connect"
	matchingv1connect "github.com/shwatanap/connect-tic-tac-toe/src/api/matching/v1/matchingv1connect"
	handler "github.com/shwatanap/connect-tic-tac-toe/src/handler/server"
)

func main() {
	// port := 50051
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// server := grpc.NewServer()

	// pb.RegisterMatchingServiceServer(server, handler.NewMatchingHandler())
	// pb.RegisterGameServiceServer(server, handler.NewGameHandler(gu))
	mux := http.NewServeMux()
	mux.Handle(gamev1connect.NewGameServiceHandler(handler.NewGameHandler()))
	mux.Handle(matchingv1connect.NewMatchingServiceHandler(handler.NewMatchingHandler()))

	http.ListenAndServe(":50051", h2c.NewHandler(mux, &http2.Server{}))
	// http.ListenAndServe(":50051", mux)

	// reflection.Register(server)

	// go func() {
	// 	log.Printf("start gRPC server port: %v", port)
	// 	server.Serve(lis)
	// }()

	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// log.Println("stopping gRPC server...")
	// server.GracefulStop()
}
