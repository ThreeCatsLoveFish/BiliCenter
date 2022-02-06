package push

var EndpointList []Endpoint

type Endpoint struct {
	// 推送URL
	url string
	// 推送token
	token string
}

func (push Endpoint) Submit(data string) {

}
