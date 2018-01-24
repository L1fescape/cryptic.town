package rpc

import (
  "context"

  db "github.com/l1fescape/cryptic.town/db"
  pb "github.com/l1fescape/cryptic.town/rpc/out"
)

var Prefix = pb.CrypticTownPathPrefix

type RPCServer struct {
  store *db.Store
}

func (s *RPCServer) SetHome(ctx context.Context, req *pb.SetHomeRequest) (*pb.Home, error) {
  home, err := s.store.SetHome(req.Token, req.Name, req.Body)
  if err != nil {
    return nil, err
  }

  return &pb.Home{Name: home.Name, Body: home.Body}, nil
}

func GetHandler(store *db.Store) pb.TwirpServer {
  return pb.NewCrypticTownServer(&RPCServer{ store: store }, nil)
}
