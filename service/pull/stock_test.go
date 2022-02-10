package pull

import "testing"

func TestEastMoneySubmit(t *testing.T) {
	pull := EastMoneyPull{}
	pull.Obtain()
}
