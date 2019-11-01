// all packages that are created specifically for this project
package dummy

type Dummy interface {
  GetAttr() string
}

type dummy struct {
  attr string
}

type DummyConfig struct {
  Attr string
}

func NewDummy(cfg DummyConfig) (Dummy, error) {
  return &dummy{
    attr: cfg.Attr,
  }, nil
}

func (du *dummy) GetAttr() string {
  return du.attr
}
