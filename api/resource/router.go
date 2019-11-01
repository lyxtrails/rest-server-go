// router file contains only setRouter function which chains all middlewares
// and set them to a given router
package resource

import (
  "github.com/gorilla/mux"
  "github.com/justinas/alice"

  "github.com/lyxtrails/rest-server-go/dep"
)

func SetRouter(router *mux.Router, d *dep.Dep, cfg ResourceHandlerConfig) {
  re, _ := NewResourceHanlder()

  router.Handle("/resources", alice.New().
    Then(re.CreateResource())).Methods("POST")
}
