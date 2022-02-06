package push

import (
	"fmt"
	"testing"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

func TestLoadConfig(t *testing.T) {
	config.WithOptions(func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
	})	

	// add driver for support yaml content
	config.AddDriver(toml.Driver)

	err := config.LoadFiles("../config/push.toml")
	if err != nil {
		panic(err)
	}

	size := config.Get("global.size")
	endpoints := make([]Endpoint, size.(int64))
	config.BindStruct("endpoints", &endpoints)

	url := fmt.Sprintf(endpoints[0].URL, endpoints[0].Token)
	if url != "https://just.for.your.test" {
		t.Fatal("BindStruct() error!")
	}
}
