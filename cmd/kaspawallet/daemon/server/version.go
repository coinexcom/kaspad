package server

import (
	"context"
	"github.com/coinexcom/kaspad/cmd/kaspawallet/daemon/pb"
	"github.com/coinexcom/kaspad/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
