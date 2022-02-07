package push

import "testing"

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
	push := NewPush(1)
	push.SetTitle("# Test turbo")
	push.SetContent("Success if you can see this info!")
	push.SetChannel([]int64{ChannelPushDeer})
	if err := push.Submit(); err != nil {
		t.Fatalf("Submit failed, error: %v", err)
	}
}
