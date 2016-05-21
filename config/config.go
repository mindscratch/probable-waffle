package config

import "flag"

var (
	crate string
)

type Config struct {
	Crate string
}

func init() {
	flag.StringVar(&crate, "crate", "http://localhost:4200", "Crate REST API URL")
}

func New() (Config, error) {
	return Config{
		Crate: crate,
	}, nil
}
