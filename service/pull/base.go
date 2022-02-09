package pull

import "time"

var pullList []Pull

type Pull interface {
	Obtain() (string, string, error)
}

func NewPull(pullId int) Pull {
	if pullId >= len(pullList) {
		return RawPull{}
	}
	return pullList[pullId]
}

type RawPull struct{}

func (RawPull) Obtain() (string, string, error) {
	return "# Heartbeat", time.Now().Format(time.RFC1123) + " from SubCenter", nil
}
