package web

import (
  "net/http"
  "fmt"
  "encoding/json"

  "github.com/gorilla/mux"
  db "github.com/l1fescape/cryptic.town/db"
)

var Prefix = "/"

var dist = "web/dist"

func GetHandler(store *db.Store) *mux.Router {
  r := mux.NewRouter()

  // Serve index file
  r.Handle(Prefix, http.FileServer(http.Dir(dist)))

  r.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
    users, err := store.GetUsers()
    if err != nil {
      users = []string{}
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
  })

  // Serve assets
  r.Handle("/fonts/font-awesome/css/{rest}", http.FileServer(http.Dir(dist)))
  r.Handle("/assets/{rest}", http.FileServer(http.Dir(dist)))

  // todo: handle /andrew/some/path/with/slashes
  r.HandleFunc("/{name}", func(w http.ResponseWriter, req *http.Request) {
    name := mux.Vars(req)["name"]
    home, err := store.GetUserHome(name)
    var body string
    if err != nil {
      w.WriteHeader(http.StatusNotFound)
      body = "no user"
    } else {
      body = home.Body
    }
    fmt.Fprint(w, "<!doctype html>" + body)
  })


  return r
}

