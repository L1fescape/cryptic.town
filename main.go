package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis"
	rpc "github.com/l1fescape/cryptic.town/rpc"
	web "github.com/l1fescape/cryptic.town/web"
	util "github.com/l1fescape/cryptic.town/util"
)

var DEFAULT_PORT = "8080"

func main() {
  // setip redis
  client := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
  })

  // Catch interrupts
  stop := make(chan os.Signal, 2)
  signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

  // Setup router
  mux := http.NewServeMux()
  mux.Handle(rpc.Prefix, rpc.GetHandler(client))
  mux.Handle(web.Prefix, web.GetHandler(client))

  // Setup server
  port := ":" + os.Getenv("PORT")
  if port == ":" {
    port = ":" + DEFAULT_PORT
  }

  h := &http.Server{
    Addr:    port,
    Handler: util.WrapHandlerWithLogging(mux),
  }

  // Run server
  go func() {
    util.Log.Printf("Listening on %s\n", port)

    if err := h.ListenAndServe(); err != nil {
      util.Log.Fatal(err)
    }
  }()

  <-stop
  util.Log.Println("\nShutting down the server...")
  ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
  h.Shutdown(ctx)
  util.Log.Println("Server gracefully stopped")
  os.Exit(2)
}
