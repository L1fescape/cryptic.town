package rpc

import (
  "context"

  pb "cryptic.town/rpc/out"
)

type RPCServer struct {}

func (s *RPCServer) MakeHome(ctx context.Context, req *pb.User) (*pb.Home, error) {
  return &pb.Home{Name: req.Name, Body: "Hi "+req.Name}, nil
}

var Prefix = pb.CrypticTownPathPrefix
var Handler = pb.NewCrypticTownServer(&RPCServer{}, nil)
