package pull

import (
	"subcenter/domain/push"
	"time"
)

var (
	location *time.Location
	pullMap  map[string]Pull
)

const (
	HeartBeat   = "HeartBeat"
	MedalHelper = "MedalHelper"
)

func init() {
	var err error
	location, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic("load location error")
	}
	addPull(HeartBeat, HeartBeatPull{})
	addPull(MedalHelper, MedalPull{})
}

type Pull interface {
	Obtain() ([]push.Data, error)
}

func addPull(name string, pull Pull) {
	if pullMap == nil {
		pullMap = make(map[string]Pull)
	}
	pullMap[name] = pull
}

func NewPull(name string) Pull {
	if pull, ok := pullMap[name]; ok {
		return pull
	} else {
		return HeartBeatPull{}
	}
}

type HeartBeatPull struct{}

func (HeartBeatPull) Obtain() ([]push.Data, error) {
	return []push.Data{
		{
			Title:   "# Heartbeat",
			Content: time.Now().In(location).Format(time.RFC1123Z),
		},
	}, nil
}
