package main

import (
	"subcenter/push"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

func init() {
	initConfig()
	push.LoadEndpoints()
}

func initConfig() {
	config.WithOptions(func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
	})
	config.AddDriver(toml.Driver)
	err := config.LoadFiles("config/push.toml")
	if err != nil {
		panic(err)
	}
}

func main() {
	println("Hello world!")
}
