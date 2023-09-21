package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	gamev1connect "github.com/shwatanap/connect-tic-tac-toe/src/connect/api/game/v1/gamev1connect"
	matchingv1connect "github.com/shwatanap/connect-tic-tac-toe/src/connect/api/matching/v1/matchingv1connect"
	handler "github.com/shwatanap/connect-tic-tac-toe/src/connect/handler/server"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle(gamev1connect.NewGameServiceHandler(handler.NewGameHandler()))
	mux.Handle(matchingv1connect.NewMatchingServiceHandler(handler.NewMatchingHandler()))

	http.ListenAndServe(":50051", h2c.NewHandler(mux, &http2.Server{}))
}
