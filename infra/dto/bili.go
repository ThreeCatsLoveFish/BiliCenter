package dto

// BiliBaseResp is basic response body of all bilibili API
type BiliBaseResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Msg     string `json:"msg,omitempty"`
}

// User info of bilibili
type User struct {
	Uid   int32  `json:"uid"`
	Uname string `json:"uname"`
	Level int32  `json:"level"`
	Color int32  `json:"color"`
}

// Anchor represents live anchor lottery
type Anchor struct {
	Id           int32  `json:"id"`
	RoomId       int32  `json:"room_id"`
	Status       int32  `json:"status"`
	AwardName    string `json:"award_name"`
	AwardNum     int32  `json:"award_num"`
	AwardImage   string `json:"award_image"`
	Danmu        string `json:"danmu"`
	Time         int32  `json:"time"`
	CurrentTime  int32  `json:"current_time"`
	JoinType     int32  `json:"join_type"`
	RequireType  int32  `json:"require_type"`
	RequireValue int32  `json:"require_value"`
	RequireText  string `json:"require_text"`
	GiftId       int32  `json:"gift_id"`
	GiftName     string `json:"gift_name"`
	GiftNum      int32  `json:"gift_num"`
	GiftPrice    int32  `json:"gift_price"`
	CurGiftNum   int32  `json:"cur_gift_num"`
	GoawayTime   int32  `json:"goaway_time"`
	AwardUsers   []User `json:"award_users"`
	ShowPanel    int32  `json:"show_panel"`
	LotStatus    int32  `json:"lot_status"`
	SendGift     int32  `json:"send_gift_ensure"`
	GoodsId      int32  `json:"goods_id"`
}

// RedPocket represents red pocket lottery
type RedPocket struct {
	RoomUid   int32 `json:"ruid"`
	RoomId    int32 `json:"room_id"`
	LotteryId int32 `json:"lot_id"`
}

// BiliAnchor is response body of anchor info
type BiliAnchor struct {
	BiliBaseResp
	Data Anchor `json:"data"`
}

type BiliNewTag struct {
	BiliBaseResp
	Data struct {
		TagId int32 `json:"tagid"`
	} `json:"data"`
}

type BiliListTag struct {
	BiliBaseResp
	Data []struct {
		Name  string `json:"name"`
		TagId int32  `json:"tagid"`
		Count int32  `json:"count"`
		Tip   string `json:"tip"`
	} `json:"data"`
}

type BiliRelation struct {
	BiliBaseResp
	Data []struct {
		Mid int32 `json:"mid"`
	} `json:"data"`
}
