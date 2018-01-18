package web

import (
  "net/http"
)


type Server struct {}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("hi!"))
}

var Prefix = "/"
var Handler = &Server{}
