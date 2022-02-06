package config

import (
    "github.com/gookit/config/v2"
    "github.com/gookit/config/v2/toml"
)

func LoadConfig() {
	config.WithOptions(func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
	})	

	// add driver for support yaml content
	config.AddDriver(toml.Driver)

	err := config.LoadFiles("config.toml")
	if err != nil {
		panic(err)
	}
}
