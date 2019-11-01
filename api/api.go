// api package is used to connect router with the routes of all resources
package api

import (
  "github.com/gorilla/mux"

  "github.com/lyxtrails/rest-server-go/api/resource"
  "github.com/lyxtrails/rest-server-go/dep"
  "github.com/lyxtrails/rest-server-go/svrcfg"
)

func SetRouter(router *mux.Router, d *dep.Dep, cfg *svrcfg.Config) {
  resource.SetRouter(router, d, resource.ResourceHandlerConfig{})
}
