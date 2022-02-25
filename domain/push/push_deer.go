package push

import (
	"net/url"
	"subcenter/infra"
)

const PushDeerName = "push_deer"

// PushDeer push
type PushDeerPush struct {
	Endpoint
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

func (push PushDeerPush) Info() Endpoint {
	return push.Endpoint
}
