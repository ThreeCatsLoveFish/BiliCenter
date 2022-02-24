package push

import (
	"fmt"
	"net/url"
	"subcenter/infra"
)

const TurboName = "turbo"

// Server-Turbo push
type TurboPush struct {
	Endpoint
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

func (push TurboPush) Info() map[string]string{
	return map[string]string{
		"Name": push.Name,
		"Type:": push.Type,
		"Url": push.URL,
		"Token": push.Token,
	}
}
