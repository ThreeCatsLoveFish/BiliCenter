package pull

import "time"

var (
	loc      *time.Location
	pullList []Pull
)

func init() {
	var err error
	loc, err = time.LoadLocation("China/Shanghai")
	if err != nil {
		panic("load location error")
	}
}

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
	return "# Heartbeat", time.Now().In(loc).Format(time.RFC1123Z), nil
}
