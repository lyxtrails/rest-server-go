// Interface between server and application
// All environment variables are imported here
package svrcfg

import (
  "github.com/caarlos0/env"
)

type Config struct {
  Port string `env:"PORT" envDefault:"8080"`
}

func GetConfig() (*Config, error) {
  cfg := new(Config)
  if err := env.Parse(cfg); err != nil {
    return cfg, err
  }
  return cfg, nil
}
