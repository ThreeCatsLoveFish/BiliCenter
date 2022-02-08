package pull

var pullList []Pull

type Pull interface {
	Obtain() (string, string, error)
}

func NewPull(pullId int64) Pull {
	return pullList[pullId]
}
