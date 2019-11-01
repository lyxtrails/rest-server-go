// main file to start the server
package main

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"

  "github.com/lyxtrails/rest-server-go/api"
  "github.com/lyxtrails/rest-server-go/dep"
  "github.com/lyxtrails/rest-server-go/svrcfg"
)

func main() {
  // Initialize server configs
  log.Println("Initialize server configs")
  cfg, err := svrcfg.GetConfig()
  if err != nil {
    panic(err)
  }

  // Initialize dependencies
  log.Println("Initialize dependencies")
  d, err := dep.NewDep(cfg)
  if err != nil {
    panic(err)
  }

  router := mux.NewRouter()

  // Set routes
  api.SetRouter(router, d, cfg)

  log.Println("Starting server at port " + cfg.Port)
  log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
