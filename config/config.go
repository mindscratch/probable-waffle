package config

import "flag"

var (
	crate  string
	output string
)

type Config struct {
	Crate  string
	Output string
}

func init() {
	flag.StringVar(&crate, "crate", "http://localhost:4200", "Crate REST API URL")
	flag.StringVar(&output, "output", "", "write output to file instead of stdout")
}

func New() (Config, error) {
	return Config{
		Crate:  crate,
		Output: output,
	}, nil
}
