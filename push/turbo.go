package push

import (
	"encoding/json"
	"fmt"
)

const (
	// 方糖服务号
	ChannelWeChatFT int64 = 9
	// PushDeer
	ChannelPushDeer int64 = 18
)

type TurboData struct {
	Title   string `json:"title"`
	Desp    string `json:"desp"`
	Channel string `json:"channel"`
}

// Create a new turbo data
func NewTurboData(title, desp string, channels []int64) TurboData {
	var channel string
	for i, ch := range channels {
		if i == 0 {
			channel += fmt.Sprintf("%d", ch)
		} else {
			channel += fmt.Sprintf("|%d", ch)
		}
	}
	return TurboData{title, desp, channel}
}

// Obtain the json string of data
func (data TurboData) ToString() string {
	body, err := json.Marshal(data)
	if err != nil {
		println("Marshal failed, error: %v", err)
		// FIXME: handle error here
		return ""
	}
	return string(body)
}
