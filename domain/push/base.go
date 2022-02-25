package push

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

var pushMap map[string]Push

func init() {
	initPush()
}

// initPush bind endpoints with config file
func initPush() {
	conf := config.NewWithOptions("push", func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
		opt.ParseEnv = true
	})
	conf.AddDriver(toml.Driver)
	err := conf.LoadFiles("config/push.toml")
	if err != nil {
		panic(err)
	}

	// Load config file
	var endpoints []Endpoint
	conf.BindStruct("endpoints", &endpoints)

	// Load token or key here
	for _, endpoint := range endpoints {
		SetEndpoint(endpoint)
	}
}

// Endpoint represents a kind of subscription
type Endpoint struct {
	Name  string `config:"name"  json:"name"`
	Type  string `config:"type"  json:"type"`
	URL   string `config:"url"   json:"url"`
	Token string `config:"token" json:"token"`
}

func SetEndpoint(endpoint Endpoint) {
	switch endpoint.Type {
	case TurboName:
		addPush(endpoint.Name, TurboPush{endpoint})
	case PushDeerName:
		addPush(endpoint.Name, PushDeerPush{endpoint})
	case PushPlusName:
		addPush(endpoint.Name, PushPlusPush{endpoint})
	}
}

func GetEndpoint() []Endpoint {
	endpoints := make([]Endpoint, 0)
	for _, push := range pushMap {
		endpoints = append(endpoints, push.Info())
	}
	return endpoints
}

// Data represents data needed for push
type Data struct {
	Title   string
	Content string
}

// Push contain all info needed for push action
type Push interface {
	Info() Endpoint
	Submit(data Data) error
}

func addPush(name string, push Push) {
	if pushMap == nil {
		pushMap = make(map[string]Push)
	}
	pushMap[name] = push
}

func NewPush(name string) Push {
	if push, ok := pushMap[name]; ok {
		return push
	}
	panic("push not found")
}
