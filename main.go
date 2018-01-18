package main

import (
  "context"
  "net/http"

  pb "cryptic.town/rpc/out"
)

type Server struct {}

func (s *Server) MakeHome(ctx context.Context, req *pb.User) (*pb.Home, error) {
  return &pb.Home{Name: req.Name, Body: "Hi "+req.Name}, nil
}

func main() {
  twirpHandler := pb.NewCrypticTownServer(&Server{}, nil)
  mux := http.NewServeMux()
  mux.Handle(pb.CrypticTownPathPrefix, twirpHandler)
  http.ListenAndServe(":8081", mux)
}
