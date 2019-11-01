// file contains definitions of resource
package resource

type Resource struct{}
type ResourceHandler struct{}
type ResourceHandlerConfig struct{}

func NewResourceHanlder() (*ResourceHandler, error) {
  return &ResourceHandler{}, nil
}
