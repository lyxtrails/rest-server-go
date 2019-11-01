// file contains all middlewares that are writen sepcifically for
// creating resource
package resource

import (
  "net/http"
)

func (re *ResourceHandler) CreateResource() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write([]byte(`{ "Response": "Resource has been successfully created" }`))
  })
}
