package pull

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

type rawPull struct {}

func (rawPull) Obtain() (string, string, error) {
	return "# Empty info", "EMPTY", nil
}
