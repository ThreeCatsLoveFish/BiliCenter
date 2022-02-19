package push

import (
	"fmt"
	"net/url"
	"subcenter/infra"
)

const (
	TurboName = "turbo"
	TurboEnv  = "TURBO"
)

// Server-Turbo push
type TurboPush struct {
	endpoint
}

// Submit data to endpoint and finish one task
func (push TurboPush) Submit(pd Data) error {
	// Prepare content and header
	rawUrl := fmt.Sprintf(push.URL, push.Token)
	data := url.Values{
		"title": []string{pd.Title},
		"desp":  []string{pd.Content},
	}
	// Submit info
	return infra.PostForm(rawUrl, data)
}
