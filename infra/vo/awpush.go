package vo

import (
	"regexp"
)

// LiveFilter store rooms and words with which lottery won't be join
type LiveFilter struct {
	Rooms []int32  `config:"rooms"`
	Words []string `config:"words"`

	WordsPat []*regexp.Regexp
}

// BiliConfig store all token and cookies used for awpush and bili live
type BiliConfig struct {
	Uid     string     `config:"uid"`
	Token   string     `config:"token"`
	Wss     string     `config:"wss"`
	Cookies []string   `config:"cookies"`
	Push    []string   `config:"push"`
	Filter  LiveFilter `config:"filter"`

	UidList []int32
}
