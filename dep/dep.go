// centralize all dependencies
package dep

import (
  "github.com/lyxtrails/rest-server-go/pkg/dummy"
  "github.com/lyxtrails/rest-server-go/svrcfg"
)

type Dep struct {
  Dummy dummy.Dummy
}

func NewDep(cfg *svrcfg.Config) (*Dep, error) {
  du, _ := dummy.NewDummy(dummy.DummyConfig{
    Attr: "real",
  })
  dep := &Dep{
    Dummy: du,
  }
  return dep, nil
}
