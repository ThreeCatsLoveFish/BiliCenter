package conf

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

var BiliConf BiliConfig

type User struct {
	Uid    int
	Csrf   string
	Buvid  string
	Cookie string `config:"cookie"`
	Push   string `config:"push"`
	Login  bool
}

// LiveFilter store rooms and words with which lottery won't be join
type LiveFilter struct {
	Rooms    []int    `config:"rooms"`
	Words    []string `config:"words"`
	WordsPat []*regexp.Regexp
}

// BiliConfig store all token and cookies used for awpush and bili live
type BiliConfig struct {
	Uid    string     `config:"uid"`
	Token  string     `config:"token"`
	Wss    string     `config:"wss"`
	Users  []User     `config:"users"`
	Filter LiveFilter `config:"filter"`
}

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
	conf.BindStruct("awpush", &BiliConf)
	BiliConf.Filter.WordsPat = make(
		[]*regexp.Regexp,
		len(BiliConf.Filter.Words),
	)
	for idx, word := range BiliConf.Filter.Words {
		pat := regexp.MustCompile(word)
		BiliConf.Filter.WordsPat[idx] = pat
	}
	for idx, user := range BiliConf.Users {
		req := http.Request{Header: map[string][]string{}}
		req.Header.Set("cookie", user.Cookie)
		uid, err := req.Cookie("DedeUserID")
		if err != nil {
			panic("cookie error")
		}
		csrf, err := req.Cookie("bili_jct")
		if err != nil {
			panic("cookie error")
		}
		buvid, err := req.Cookie("LIVE_BUVID")
		if err != nil {
			panic("cookie error")
		}
		num, err := strconv.ParseInt(uid.Value, 10, 32)
		if err != nil {
			panic("uid parse err")
		}
		BiliConf.Users[idx].Uid = int(num)
		BiliConf.Users[idx].Csrf = csrf.Value
		BiliConf.Users[idx].Buvid = buvid.Value
		BiliConf.Users[idx].Login = true
	}
}
