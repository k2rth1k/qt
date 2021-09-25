package api

import (
	"github.com/k2rth1k/qt/pkg/config"
	"github.com/k2rth1k/qt/pkg/db"
	qt "github.com/k2rth1k/qt/pkg/proto"
	log2 "github.com/k2rth1k/qt/utilities/log"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type QuickTradeService struct {
	store  db.Store
	logger *zap.SugaredLogger
}

func NewQuickTradeService() (*QuickTradeService, error) {
	cfg := config.GetServiceConfig()
	store, err := db.NewSQL(cfg.DBConfig)
	if err != nil {
		log.Fatal("Failed to create DB connection due to following error:", err)
		return nil, err
	}
	service := &QuickTradeService{store: store}
	logger := log2.InitZapLog()
	service.logger = logger
	return service, nil
}

func NewServer(grpcEndpoint string, interceptor grpc.ServerOption) {
	service, err := NewQuickTradeService()
	if err != nil {
		log.Fatal("NewQuickTradeService has failed")
	}
	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer(interceptor)
	qt.RegisterQuickTradeServer(s, service)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
