// mock dependency for testing
package mock

type MockDummy struct {
  attr string
}

func NewMockDummy() (*MockDummy, error) {
  return &MockDummy{
    attr: "mock",
  }, nil
}

func (du *MockDummy) GetAttr() string {
  return du.attr
}
