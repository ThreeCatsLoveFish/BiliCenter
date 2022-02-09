package push

import (
	"fmt"
	"net/url"
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
func (push turboPush) SetTitle(title string) turboPush {
	push.title = title
	return push
}

// Set body of data
func (push turboPush) SetContent(content string) turboPush {
	push.desp = content
	return push
}

// ToValue gives a map with value embedded
func (push turboPush) ToValue() url.Values {
	return url.Values{
		"title": []string{push.title},
		"desp":  []string{push.desp},
	}
}

// Submit data to endpoint and finish one task
func (push turboPush) Submit(title, content string) error {
	// Prepare content and header
	url := fmt.Sprintf(push.URL, push.Token)
	data := push.SetTitle(title).SetContent(content).ToValue()
	// Submit info
	return manager.PostForm(url, data)
}
