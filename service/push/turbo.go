package push

import (
	"fmt"
)

const (
	// Data name
	TurboName = "turbo"
	// Token name
	TurboEnv = "TURBO"

	// FangTang WeChat
	ChannelWeChatFT int64 = 9
	// PushDeer
	ChannelPushDeer int64 = 18
)

func init() {
	registerData(TurboName, &TurboData{})
}

// Server-Turbo data type
type TurboData struct {
	title   string
	desp    string
	channel string
}

// Set title of data
func (TurboData) DataName() string {
	return TurboName
}

// Set title of data
func (data *TurboData) SetTitle(title string) {
	data.title = title
}

// Set body of data
func (data *TurboData) SetContent(content string) {
	data.desp = content
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
	data.channel = channel
}

// Marshal the data and obtain json string
func (data *TurboData) ToString() string {
	return fmt.Sprintf("title=%s&desp=%s&channel=%s",
		data.title,
		data.desp,
		data.channel,
	)
}
