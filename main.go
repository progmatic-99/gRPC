package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/progmatic-99/gRPC/server"
	"google.golang.org/grpc"

	protos "github.com/progmatic-99/gRPC/proto/currency"
)

func main() {
	log := hclog.Default()

	cs := server.NewCurrency(log)
	gs := grpc.NewServer()

	protos.RegisterCurrencyServer(gs, cs)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error("Unable to listen.", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
