package server

import (
	"context"

	"github.com/solethus/dealership-service/internal/service"
	pb "github.com/solethus/shared-proto/proto/dealership"
)

type DealershipServer struct {
	pb.UnimplementedDealershipServiceServer
	service service.DealershipService
}

func NewDealershipServer(service service.DealershipService) *DealershipServer {
	return &DealershipServer{
		service: service,
	}
}

func (s *DealershipServer) GetDealershipInfo(ctx context.Context, req *pb.GetDealershipInfoRequest) (*pb.DealershipInfo, error) {
	//TODO implement me
	panic("implement me")
}
