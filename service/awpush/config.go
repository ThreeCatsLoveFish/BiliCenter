package awpush

import (
	"regexp"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

var biliConfig BiliConfig

func init() {
	initAWPush()
}

// initAWPush load awpush and bili config
func initAWPush() {
	conf := config.NewWithOptions("bili", func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
	})
	conf.AddDriver(toml.Driver)
	err := conf.LoadFiles("config/bili.toml")
	if err != nil {
		panic(err)
	}
	conf.BindStruct("awpush", &biliConfig)
	biliConfig.Filter.WordsPat = make([]*regexp.Regexp, len(biliConfig.Filter.Words))
	for idx, word := range biliConfig.Filter.Words {
		pat := regexp.MustCompile(word)
		biliConfig.Filter.WordsPat[idx] = pat
	}
}
