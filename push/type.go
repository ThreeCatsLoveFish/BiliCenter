package push

import (
	"fmt"
	"subcenter/manager"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

var (
	endpoints []Endpoint
	dataMap   map[string]Data
)

func init() {
	loadEndpoints()
}

// Endpoint represents a kind of subscription
type Endpoint struct {
	Type  string `config:"type"`
	URL   string `config:"url"`
	Token string `config:"token"`
}

func loadEndpoints() {
	pushConf := config.NewWithOptions("push", func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
	})
	pushConf.AddDriver(toml.Driver)
	err := pushConf.LoadFiles("../config/push.toml")
	if err != nil {
		panic(err)
	}
	size := pushConf.Get("global.size")
	endpoints = make([]Endpoint, size.(int64))
	pushConf.BindStruct("endpoints", &endpoints)
}

// Data contain all info needed for push
type Data interface {
	DataName() string
	SetTitle(title string)
	SetContent(body string)
	SetChannel(channel []int64)
	ToString() string
}

// Default Data type
type DefaultData struct{}

func (DefaultData) DataName() string           { return "" }
func (DefaultData) ToString() string           { return "" }
func (DefaultData) SetTitle(title string)      {}
func (DefaultData) SetContent(body string)     {}
func (DefaultData) SetChannel(channel []int64) {}

// Register data with the specific name
func registerData(name string, data Data) {
	if dataMap == nil {
		dataMap = make(map[string]Data)
	}
	dataMap[name] = data
}

func getData(name string) Data {
	data, ok := dataMap[name]
	if !ok {
		return &DefaultData{}
	}
	return data
}

type Push struct {
	Endpoint
	Data
}

func NewPush(id int64) Push {
	endpoint := endpoints[id]
	data := getData(endpoint.Type)
	return Push{endpoint, data}
}

// Submit the data to endpoint and finish one push task
func (push Push) Submit() error {
	url := fmt.Sprintf(push.URL, push.Token)
	data := push.ToString()
	return manager.Post(url, data)
}
