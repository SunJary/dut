package pkg

type Config struct {
	PrintLevel   int
	ByteSzieUnit string
	Usage        bool
	Interact     bool
}

func (c *Config) Init() {
	if c.PrintLevel >= 1 {
		c.PrintLevel--
	}
}
