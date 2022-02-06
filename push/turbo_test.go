package push

import (
	"testing"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

func TestTurboData(t *testing.T) {
	data := TurboData{}
	data.SetTitle("# Test function")
	data.SetContent("Test ONLY")
	data.SetChannel([]int64{ChannelPushDeer, ChannelWeChatFT})

	result := data.ToString()
	except := `title=# Test function&desp=Test ONLY&channel=18|9`
	if result != except {
		t.Fatalf("ToString() error! get: %s, except: %s", result, except)
	}
}

func TestTurboPush(t *testing.T) {
	config.WithOptions(func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
	})
	config.AddDriver(toml.Driver)
	err := config.LoadFiles("../config/push.toml")
	if err != nil {
		t.Log(err)
	}
	LoadEndpoints()

	push := NewPush(1)
	push.SetTitle("# Test turbo")
	push.SetContent("Success if you can see this info!")
	push.SetChannel([]int64{ChannelPushDeer, ChannelWeChatFT})
	push.Submit()
}
