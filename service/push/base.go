package push

import (
	"fmt"
	"subcenter/manager"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

var (
	endpoints []Endpoint
	dataMap   map[string]PushData
)

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
	endpoints = make([]Endpoint, size)
	pushConf.BindStruct("endpoints", &endpoints)

	// Load token or key here
	pushConf.LoadOSEnv([]string{TurboEnv}, false)
	// TODO: support multi tokens
	turbo := pushConf.Get(TurboEnv).(string)
	for idx, endpoint := range endpoints {
		if endpoint.Type == TurboName {
			endpoints[idx].Token = turbo
		}
	}
}

// PushData contain all info needed for push action
type PushData interface {
	DataName() string
	SetTitle(title string)
	SetContent(body string)
	SetChannel(channel []int64)
	ToString() string
}

type EmptyPushData struct{}

func (EmptyPushData) DataName() string           { return "" }
func (EmptyPushData) ToString() string           { return "" }
func (EmptyPushData) SetTitle(title string)      {}
func (EmptyPushData) SetContent(body string)     {}
func (EmptyPushData) SetChannel(channel []int64) {}

// Register data with the specific name
func registerData(name string, data PushData) {
	if dataMap == nil {
		dataMap = make(map[string]PushData)
	}
	dataMap[name] = data
}

// Take out data with the specific name
func getData(name string) PushData {
	data, ok := dataMap[name]
	if !ok {
		return &EmptyPushData{}
	}
	return data
}

type Push struct {
	Endpoint
	PushData
}

func NewPush(id int64) Push {
	endpoint := endpoints[id]
	data := getData(endpoint.Type)
	return Push{endpoint, data}
}

// Submit data to endpoint and finish one task
func (push Push) Submit() error {
	url := fmt.Sprintf(push.URL, push.Token)
	data := push.ToString()
	return manager.Post(url, data)
}
