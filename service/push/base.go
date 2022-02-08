package push

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

var pushList []Push

func init() {
	initPush()
}

// endpoint represents a kind of subscription
type endpoint struct {
	Type  string `config:"type"`
	URL   string `config:"url"`
	Token string `config:"token"`
}

// initPush bind endpoints with config file
func initPush() {
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
	endpoints := make([]endpoint, size)
	pushConf.BindStruct("endpoints", &endpoints)

	// Load token or key here
	pushConf.LoadOSEnv([]string{TurboEnv}, false)
	// TODO: support multi tokens
	turbo := pushConf.Get(TurboEnv).(string)
	for _, endpoint := range endpoints {
		if endpoint.Type == TurboName {
			endpoint.Token = turbo
			pushList = append(pushList, &turboPush{
				endpoint: endpoint,
			})
		}
	}
}

// Push contain all info needed for push action
type Push interface {
	Submit(title, content string) error
}

func NewPush(pushId int) Push {
	if pushId >= len(pushList) {
		return rawPush{}
	}
	return pushList[pushId]
}

type rawPush struct {
	endpoint
}

func (rawPush) Submit(title, content string) error {
	return nil
}
