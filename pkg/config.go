package pkg

import "strings"

type Config struct {
	PrintLevel   int
	ByteSzieUnit string
	Usage        bool
	Interact     bool
}

func (c *Config) Init() error {
	if c.PrintLevel >= 1 {
		c.PrintLevel--
	}

	c.ByteSzieUnit = strings.ToUpper(c.ByteSzieUnit)
	err := IsAllowUnit(*c)

	return err
}
