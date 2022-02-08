package push

import (
	"fmt"
	"subcenter/manager"
)

const (
	// Data name
	TurboName = "turbo"
	// Token env name
	TurboEnv = "TURBO"
)

// Server-Turbo push
type TurboPush struct {
	Endpoint
	title string
	desp  string
}

// Name of push
func (TurboPush) PushName() string {
	return TurboName
}

// Set title of data
func (push *TurboPush) SetTitle(title string) {
	push.title = title
}

// Set body of data
func (push *TurboPush) SetContent(content string) {
	push.desp = content
}

// Marshal the data and obtain json string
func (push *TurboPush) ToString() string {
	return fmt.Sprintf("title=%s&desp=%s",
		push.title,
		push.desp,
	)
}

// Submit data to endpoint and finish one task
func (push *TurboPush) Submit(title, content string) error {
	// Prepare content and header
	push.SetTitle(title)
	push.SetContent(content)

	// Submit info
	url := fmt.Sprintf(push.URL, push.Token)
	data := push.ToString()
	return manager.Post(url, data)
}
