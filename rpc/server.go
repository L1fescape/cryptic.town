package rpc

import (
  "context"

  "github.com/go-redis/redis"
  pb "cryptic.town/rpc/out"
)

var Prefix = pb.CrypticTownPathPrefix

type RPCServer struct {
  client *redis.Client
}

func (s *RPCServer) MakeHome(ctx context.Context, req *pb.Home) (*pb.Home, error) {
  err := s.client.Set(req.Name, req.Body, 0).Err()
	if err != nil {
		panic(err)
  }

  return &pb.Home{Name: req.Name, Body: req.Body}, nil
}

func GetHandler(client *redis.Client) pb.TwirpServer {
  return pb.NewCrypticTownServer(&RPCServer{client: client}, nil)
}
