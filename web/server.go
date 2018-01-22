package web

import (
  "net/http"
  "fmt"
  "encoding/json"

  "github.com/go-redis/redis"
  "github.com/gorilla/mux"
)

var Prefix = "/"

var dist = "web/dist"

func GetHandler(client *redis.Client) *mux.Router {
  r := mux.NewRouter()

  // Serve index file
  r.Handle(Prefix, http.FileServer(http.Dir(dist)))

  r.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
    keys, _, err := client.Scan(0, "", 10).Result()
    if err != nil {
      keys = []string{}
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(keys)
  })

  // Serve assets
  r.Handle("/fonts/font-awesome/css/{rest}", http.FileServer(http.Dir(dist)))
  r.Handle("/assets/{rest}", http.FileServer(http.Dir(dist)))

  // todo: handle /andrew/some/path/with/slashes
  r.HandleFunc("/{name}", func(w http.ResponseWriter, req *http.Request) {
    name := mux.Vars(req)["name"]
    val, err := client.Get(name).Result()
    if err != nil {
      w.WriteHeader(http.StatusNotFound)
      val = "no user"
    }
    fmt.Fprint(w, "<!doctype html>" + val)
  })


  return r
}

