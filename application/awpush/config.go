package awpush

import (
	"net/http"
	"regexp"
	"strconv"
	"subcenter/infra/vo"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

var biliConfig vo.BiliConfig

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
	biliConfig.Filter.WordsPat = make(
		[]*regexp.Regexp,
		len(biliConfig.Filter.Words),
	)
	for idx, word := range biliConfig.Filter.Words {
		pat := regexp.MustCompile(word)
		biliConfig.Filter.WordsPat[idx] = pat
	}
	biliConfig.UidList = make([]int32, 0)
	for idx, cookie := range biliConfig.Cookies {
		req := http.Request{}
		req.Header.Set("cookie", cookie)
		uid, err := req.Cookie("DedeUserID")
		if err != nil {
			panic("cookie error")
		}
		num, err := strconv.ParseInt(uid.Value, 10, 32)
		if err != nil {
			panic("uid parse err")
		}
		biliConfig.UidList[idx] = int32(num)
	}
}
