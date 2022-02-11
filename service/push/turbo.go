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
}

// Submit data to endpoint and finish one task
func (push turboPush) Submit(pd Data) error {
	// Prepare content and header
	rawUrl := fmt.Sprintf(push.URL, push.Token)
	data := url.Values{
		"title": []string{pd.Title},
		"desp":  []string{pd.Content},
	}
	// Submit info
	return manager.PostForm(rawUrl, data)
}
