package push

import (
	"fmt"
)

const (
	// Data name
	dataName = "turbo"

	// FangTang WeChat
	ChannelWeChatFT int64 = 9
	// PushDeer
	ChannelPushDeer int64 = 18
)

func init() {
	registerData(dataName, &TurboData{})
}

// Server-Turbo data type
type TurboData struct {
	Title   string
	Desp    string
	Channel string
}

// Set title of data
func (TurboData) DataName() string {
	return dataName
}

// Set title of data
func (data *TurboData) SetTitle(title string) {
	data.Title = title
}

// Set body of data
func (data *TurboData) SetContent(content string) {
	data.Desp = content
}

// Set channel of data
func (data *TurboData) SetChannel(channels []int64) {
	var channel string
	for i, ch := range channels {
		if i == 0 {
			channel += fmt.Sprintf("%d", ch)
		} else {
			channel += fmt.Sprintf("|%d", ch)
		}
	}
	data.Channel = channel
}

// Marshal the data and obtain json string
func (data *TurboData) ToString() string {
	return fmt.Sprintf("title=%s&desp=%s&channel=%s",
		data.Title,
		data.Desp,
		data.Channel,
	)
}
