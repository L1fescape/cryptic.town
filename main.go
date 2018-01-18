package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "time"
  "syscall"

  pb "cryptic.town/rpc/out"
)

type Server struct {
  logger *log.Logger
  mux    *http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

type RPCServer struct {}

func (s *RPCServer) MakeHome(ctx context.Context, req *pb.User) (*pb.Home, error) {
  return &pb.Home{Name: req.Name, Body: "Hi "+req.Name}, nil
}

type WebServer struct {}

func (s *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("hi!"))
}

func main() {
  stop := make(chan os.Signal, 2)
  signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
  logger := log.New(os.Stdout, "", 0)

  // RPC Methods
  twirpHandler := pb.NewCrypticTownServer(&RPCServer{}, nil)
  mux := http.NewServeMux()
  mux.Handle(pb.CrypticTownPathPrefix, twirpHandler)

  // Serve HTML
  mux.Handle("/", &WebServer{})

  // Setup server
  port := ":" + os.Getenv("PORT")
  if port == ":" { port = ":8081" }

  s := &Server{ logger: log.New(os.Stdout, "", 0), mux: mux, }
  h := &http.Server{Addr: port, Handler: s}

  go func() {
    logger.Printf("Listening on %s\n", port)

    if err := h.ListenAndServe(); err != nil {
      logger.Fatal(err)
    }
  }()

  <-stop

  logger.Println("\nShutting down the server...")

  ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

  h.Shutdown(ctx)

  logger.Println("Server gracefully stopped")
}
