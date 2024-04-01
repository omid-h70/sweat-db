package gapi

import (
	"context"
	"omidh70/sweat-db/pb"
)

func (server *Server) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingResp, error) {
	return &pb.PingResp{
		Msg: req.GetMsg(),
	}, nil
}
