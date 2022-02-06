package push

import (
	// "fmt"

	"github.com/gookit/config/v2"
)

var EndpointMap map[int]Endpoint

// Endpoint represents a kind of subscription
type Endpoint struct {
	Id    int    `json:"id"`
	URL   string `json:"url"`
	Token string `json:"token"`
}

// Submit the data to endpoint and finish one push task
func (push Endpoint) Submit(data string) {
	// TODO: add submit function and test function
	// fmt.Sprintf(push.URL, push.Token)
}

func LoadConfig() {
	size := config.Get("size.push")
	endpoints := make([]Endpoint, size.(int64))
	config.BindStruct("endpoints", &endpoints)
	EndpointMap = make(map[int]Endpoint)
	for _, endpoint := range endpoints {
		if endpoint.Id < 0 {
			// For test only
			continue
		}
		EndpointMap[endpoint.Id] = endpoint
	}
}
