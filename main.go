package main

import (
  "context"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  db "github.com/l1fescape/cryptic.town/db"
  rpc "github.com/l1fescape/cryptic.town/rpc"
  web "github.com/l1fescape/cryptic.town/web"
  slack "github.com/l1fescape/cryptic.town/slack"
  util "github.com/l1fescape/cryptic.town/util"
)

var DEFAULT_PORT = "8080"

func main() {
  // Catch interrupts
  stop := make(chan os.Signal, 2)
  signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

  // Setup db
  store := db.New(&db.Options{
    Addr:     "localhost:6379",
    Password: "",
  })

  // Setup slack bot
  s := slack.New(os.Getenv("SLACK_TOKEN"), store)

  // Start slack
  go func() {
    s.Start()
  }()

  // Setup router
  mux := http.NewServeMux()
  mux.Handle(rpc.Prefix, rpc.GetHandler(store))
  mux.Handle(web.Prefix, web.GetHandler(store))

  // Setup web server
  port := ":" + os.Getenv("PORT")
  if port == ":" {
    port = ":" + DEFAULT_PORT
  }
  h := &http.Server{
    Addr:    port,
    Handler: util.WrapHandlerWithLogging(mux),
  }

  // Start web server
  go func() {
    util.Log.Printf("Listening on %s\n", port)

    if err := h.ListenAndServe(); err != nil {
      util.Log.Fatal(err)
    }
  }()

  <-stop

  util.Log.Println("\nShutting down web server...")
  ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
  h.Shutdown(ctx)
  util.Log.Println("Server gracefully stopped")

  util.Log.Println("Disconnecting from slack...")
  s.Stop()
  util.Log.Println("Slack disconnected")

  util.Log.Println("Disconnecting from redis...")
  store.Quit()
  util.Log.Println("Redis disconnected")
}
