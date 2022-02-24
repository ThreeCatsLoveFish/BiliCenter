package push

import (
	"net/url"
	"subcenter/infra"
)

const PushPlusName = "push_plus"

// WeChat PushPlus push
type PushPlusPush struct {
	Endpoint
}

// Submit data to endpoint and finish one task
func (push PushPlusPush) Submit(pd Data) error {
	// Prepare content and header
	data := url.Values{
		"token":    []string{push.Token},
		"title":    []string{pd.Title},
		"content":  []string{pd.Content},
		"template": []string{"markdown"},
	}
	// Submit info
	_, err := infra.GetWithParams(push.URL, data)
	return err
}

func (push PushPlusPush) Info() map[string]string{
	return map[string]string{
		"Name": push.Name,
		"Type:": push.Type,
		"Url": push.URL,
		"Token": push.Token,
	}
}