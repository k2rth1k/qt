package main

import (
	"context"
	"flag"
	"github.com/k2rth1k/qt/pkg/api"
	qt "github.com/k2rth1k/qt/pkg/proto"
	"github.com/k2rth1k/qt/utilities/authentication"
	"github.com/k2rth1k/qt/utilities/log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50444", "gRPC server endpoint")
	httpServerEndpoint = flag.String("http-server-endpoint", "localhost:50443", "http server endpoint")
	Logger             = log.InitZapLog()
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := qt.RegisterQuickTradeHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	go api.NewServer(*grpcServerEndpoint, withServerUnaryInterceptor())

	Logger.Info("starting grpcServer at port " + *grpcServerEndpoint)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	Logger.Info("starting httpServer at port " + *httpServerEndpoint)

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	err = http.ListenAndServe(*httpServerEndpoint, corsWrapper.Handler(mux))
	if err != nil {
		Logger.Error("failed to start http server due to err: ", err)
		return err
	}
	return err
}

func main() {
	flag.Parse()
	defer glog.Flush()
	if err := run(); err != nil {
		Logger.Error("failed to start server due to error: ", err)
	}
}

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(authentication.ServerInterceptor)
}
