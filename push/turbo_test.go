package push

import "testing"

func TestTurboData(t *testing.T) {
	data := TurboData{}
	data.SetTitle("# Test function")
	data.SetContent("Test ONLY")
	data.SetChannel([]int64{ChannelPushDeer, ChannelWeChatFT})

	str := data.ToString()
	correct := `{"title":"# Test function","desp":"Test ONLY","channel":"18|9"}`
	if str != correct {
		t.Fatal("ToString() error!")
	}
}
