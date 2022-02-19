package vo

import (
	"regexp"
)

type User struct {
	Uid    int32
	Cookie string `config:"cookie"`
	Push   string `config:"push"`
}

// LiveFilter store rooms and words with which lottery won't be join
type LiveFilter struct {
	Rooms    []int32  `config:"rooms"`
	Words    []string `config:"words"`
	WordsPat []*regexp.Regexp
}

// BiliConfig store all token and cookies used for awpush and bili live
type BiliConfig struct {
	Uid    string     `config:"uid"`
	Token  string     `config:"token"`
	Wss    string     `config:"wss"`
	Users  []User     `config:"users"`
	Filter LiveFilter `config:"filter"`
}
