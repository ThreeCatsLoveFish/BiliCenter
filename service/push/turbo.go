package push

import (
	"fmt"
	"subcenter/manager"
)

const (
	TurboName = "turbo"
	TurboEnv  = "TURBO"
)

// Server-Turbo push
type turboPush struct {
	endpoint
	title string
	desp  string
}

// Set title of data
func (push *turboPush) SetTitle(title string) {
	push.title = title
}

// Set body of data
func (push *turboPush) SetContent(content string) {
	push.desp = content
}

// Marshal the data and obtain json string
func (push *turboPush) ToString() string {
	return fmt.Sprintf("title=%s&desp=%s",
		push.title,
		push.desp,
	)
}

// Submit data to endpoint and finish one task
func (push *turboPush) Submit(title, content string) error {
	// Prepare content and header
	push.SetTitle(title)
	push.SetContent(content)

	// Submit info
	url := fmt.Sprintf(push.URL, push.Token)
	data := push.ToString()
	return manager.Post(url, data)
}
