package push

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

var pushList []Push

func init() {
	initEndpoint()
}

// Endpoint represents a kind of subscription
type Endpoint struct {
	Type  string `config:"type"`
	URL   string `config:"url"`
	Token string `config:"token"`
}

// initEndpoint bind endpoints with config file
func initEndpoint() {
	pushConf := config.NewWithOptions("push", func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
		opt.ParseEnv = true
	})
	pushConf.AddDriver(toml.Driver)
	err := pushConf.LoadFiles("../../config/push.toml")
	if err != nil {
		panic(err)
	}

	// Load config file
	size := pushConf.Get("global.size").(int64)
	endpoints := make([]Endpoint, size)
	pushConf.BindStruct("endpoints", &endpoints)

	// Load token or key here
	pushConf.LoadOSEnv([]string{TurboEnv}, false)
	// TODO: support multi tokens
	turbo := pushConf.Get(TurboEnv).(string)
	for _, endpoint := range endpoints {
		if endpoint.Type == TurboName {
			endpoint.Token = turbo
			pushList = append(pushList, &TurboPush{
				Endpoint: endpoint,
			})
		}
	}
}

// Push contain all info needed for push action
type Push interface {
	PushName() string
	Submit(title, content string) error
}

func NewPush(pushId int64) Push {
	return pushList[pushId]
}

type RawPush struct {
	Endpoint
}

func (RawPush) PushName() string {
	return "RawPush"
}

func (RawPush) Submit(title, content string) error {
	return nil
}
