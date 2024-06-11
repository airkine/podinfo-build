package grpc

import (
	"context"

	pb "github.com/airkine/podinfo-build/pkg/api/grpc/version"
	"github.com/airkine/podinfo-build/pkg/version"
	"go.uber.org/zap"
)

type VersionServer struct {
	pb.UnimplementedVersionServiceServer
	config *Config
	logger *zap.Logger
}

func (s *VersionServer) Version(ctx context.Context, req *pb.VersionRequest) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{Version: version.VERSION, Commit: version.REVISION}, nil
}
