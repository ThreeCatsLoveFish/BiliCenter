package push

import (
	"net/url"
	"subcenter/manager"
)

const (
	PushDeerName = "push_deer"
	PushDeerEnv  = "PUSHDEER"
)

// Server-Turbo push
type PushDeerPush struct {
	endpoint
}

// Submit data to endpoint and finish one task
func (push PushDeerPush) Submit(pd Data) error {
	// Prepare content and header
	data := url.Values{
		"pushkey": []string{push.Token},
		"text": []string{pd.Title},
		"desp":  []string{pd.Content},
	}
	// Submit info
	return manager.PostForm(push.URL, data)
}
