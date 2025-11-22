package health

import "github.com/Johnny2705/go-tunnel-gateway/internal/config"

type Checker struct {
	cfg *config.Config
}

func NewChecker(cfg *config.Config) *Checker {
	return &Checker{
		cfg: cfg,
	}
}

func (c *Checker) CheckLiveness() error {
	return nil
}

func (c *Checker) CheckReadiness() error {
	return nil
}
