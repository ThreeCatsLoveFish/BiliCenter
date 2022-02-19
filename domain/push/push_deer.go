package push

import (
	"net/url"
	"subcenter/infra"
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
		"text":    []string{pd.Title},
		"desp":    []string{pd.Content},
	}
	// Submit info
	return infra.PostForm(push.URL, data)
}
