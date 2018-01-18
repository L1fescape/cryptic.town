package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "time"
  "syscall"

  rpc "cryptic.town/rpc"
  web "cryptic.town/web"
)

type Server struct {
  logger *log.Logger
  mux    *http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func main() {
  // Setup logging and error handling
  stop := make(chan os.Signal, 2)
  signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
  logger := log.New(os.Stdout, "", 0)

  // Setup router
  mux := http.NewServeMux()
  mux.Handle(web.Prefix, web.Handler)
  mux.Handle(rpc.Prefix, rpc.Handler)

  // Setup server
  port := ":" + os.Getenv("PORT")
  if port == ":" { port = ":8081" }

  s := &Server{
    logger: log.New(os.Stdout, "", 0),
    mux: mux,
  }
  h := &http.Server{
    Addr: port,
    Handler: s,
  }

  // Run server
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
