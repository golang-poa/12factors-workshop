package main

import (
	env "github.com/Netflix/go-env"
)

type Configuration struct {
	DB string `env:"WORKSHOP_DB"`
	Listen string `env:"WORKSHOP_LISTEN"`
}

var conf Configuration

func init() {
	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		panic(err)
	}
	if conf.DB == "" {
		conf.DB = DB
	}
	if conf.Listen == "" {
		conf.Listen = ":9090"
	}
}