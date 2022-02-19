package dto

// Verify for awpush service
type Verify struct {
	Code   string `json:"code"`
	Uid    string `json:"uid"`
	Apikey string `json:"apikey"`
}

// AWPush will return message for tasks of both poll and lottery, and type of
// message can be judged by check the string value of `Type`. Each message will
// be handled by specific handler.
//
// RawMsg represents original message
type RawMsg struct {
	Code int32  `json:"code"`
	Type string `json:"type"`
}

type AreaData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Page int32  `json:"page"`
	Size int32  `json:"size"`
}

type Poll struct {
	Task      string `json:"task"`
	MaxRoom   int32  `json:"max_room"`
	SleepTime int32  `json:"sleep_time"`
	Interval  int32  `json:"interval"`
	Secret    string `json:"secret"`

	AreaData `json:"area_data,omitempty"`
}

// TaskMsg represents message contain poll task
type TaskMsg struct {
	RawMsg

	Data Poll `json:"data"`
}

// TaskMsg represents message contain poll task
type AnchorMsg struct {
	RawMsg

	Data Anchor `json:"data"`
}

// Callback is used for response awpush poll task
type Callback struct {
	Code   string `json:"code"`
	Uid    string `json:"uid"`
	Secret string `json:"secret"`
}
