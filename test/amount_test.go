package test

import (
	"regexp"
	"strconv"
	"testing"
)

func TestRedPocket(t *testing.T) {
	redPocket := regexp.MustCompile(`(([1-9][0-9]*)(\.[0-9]{1,2})?)元`)
	testCases := []struct {
		val string
	}{
		{val: "100元",},
		{val: "1.5元",},
		{val: "10元",},
		{val: "5.20元",},
		{val: "20元",},
		{val: "30元",},
		{val: "66.66元",},
	}
	for _, tC := range testCases {
		t.Run(tC.val, func(t *testing.T) {
			res := redPocket.FindSubmatch([]byte(tC.val))
			if len(res) > 0 {
				amount := string(res[1])
				money, err := strconv.ParseFloat(amount, 32)
				if err != nil {
					t.Error(err)
				}
				t.Logf("RedPocket %s, %f", res[0], money)
			}
		})
	}
}
