package awpush

import (
	"encoding/json"
	"subcenter/manager"
)

type Verify struct {
	Code   string `json:"code"`
	Uid    string `json:"uid"`
	Apikey string `json:"apikey"`
}

func NewVerify(uid, apiKey string) []byte {
	data := Verify{
		Code:   "VERIFY_APIKEY",
		Uid:    uid,
		Apikey: apiKey,
	}
	dataStr, err := json.Marshal(data)
	if err != nil {
		return []byte("")
	}
	return manager.PakoDeflate(dataStr)
}

// RawMsg represents original message
type RawMsg struct {
	Code int32       `json:"code"`
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// TaskMsg represents message contain poll task
type TaskMsg struct {
	Code int32  `json:"code"`
	Type string `json:"type"`
	Data Poll   `json:"data"`
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
type AnchorMsg struct {
	Code int32  `json:"code"`
	Type string `json:"type"`
	Data Anchor `json:"data"`
}

type Anchor struct {
	Id           int32    `json:"id"`
	RoomId       int32    `json:"room_id"`
	Status       int32    `json:"status"`
	AwardName    string   `json:"award_name"`
	AwardNum     int32    `json:"award_num"`
	AwardImage   string   `json:"award_image"`
	Barrage      string   `json:"danmu"`
	Time         int32    `json:"time"`
	CurrentTime  int32    `json:"current_time"`
	JoinType     int32    `json:"join_type"`
	RequireType  int32    `json:"require_type"`
	RequireValue int32    `json:"require_value"`
	RequireText  string   `json:"require_text"`
	GiftId       int32    `json:"gift_id"`
	GiftName     string   `json:"gift_name"`
	GiftNum      int32    `json:"gift_num"`
	GiftPrice    int32    `json:"gift_price"`
	CurGiftNum   int32    `json:"cur_gift_num"`
	GoawayTime   int32    `json:"goaway_time"`
	AwardUsers   []string `json:"award_users"`
	ShowPanel    int32    `json:"show_panel"`
	LotStatus    int32    `json:"lot_status"`
	SendGift     int32    `json:"send_gift_ensure"`
	GoodsId      int32    `json:"goods_id"`
}

type Resp struct {
	Code   string `json:"code"`
	Uid    string `json:"uid"`
	Secret string `json:"secret"`
}
