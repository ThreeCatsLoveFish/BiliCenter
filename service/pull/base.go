package pull

import "time"

var pullList []Pull

type Pull interface {
	Obtain() (string, string, error)
}

func NewPull(pullId int) Pull {
	if pullId >= len(pullList) {
		return rawPull{}
	}
	return pullList[pullId]
}

type rawPull struct{}

func (rawPull) Obtain() (string, string, error) {
	return "# Heartbeat", time.Now().Format(time.RFC1123) + " From SubCenter", nil
}
