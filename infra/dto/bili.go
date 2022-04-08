package dto

// BiliBaseResp is basic response body of all bilibili API
type BiliBaseResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Msg     string `json:"msg,omitempty"`
}

// User info of bilibili
type User struct {
	UID   int32  `json:"uid"`
	Uname string `json:"uname"`
	Level int32  `json:"level"`
	Color int32  `json:"color"`
}

// Anchor represents live anchor lottery
type Anchor struct {
	ID           int32  `json:"id"`
	RoomID       int32  `json:"room_id"`
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
	GiftID       int32  `json:"gift_id"`
	GiftName     string `json:"gift_name"`
	GiftNum      int32  `json:"gift_num"`
	GiftPrice    int32  `json:"gift_price"`
	CurGiftNum   int32  `json:"cur_gift_num"`
	GoawayTime   int32  `json:"goaway_time"`
	AwardUsers   []User `json:"award_users"`
	ShowPanel    int32  `json:"show_panel"`
	LotStatus    int32  `json:"lot_status"`
	SendGift     int32  `json:"send_gift_ensure"`
	GoodsID      int32  `json:"goods_id"`
}

// RedPocket represents red pocket lottery
type RedPocket struct {
	LotteryID       int32  `json:"lot_id"`
	SenderUID       int32  `json:"sender_uid"`
	SenderName      string `json:"sender_name"`
	JoinRequirement int32  `json:"join_requirement"`
	Danmu           string `json:"danmu"`
	CurrentTime     int32  `json:"current_time"`
	StartTime       int32  `json:"start_time"`
	EndTime         int32  `json:"end_time"`
	LastTime        int32  `json:"last_time"`
	RemoveTime      int32  `json:"remove_time"`
	ReplaceTime     int32  `json:"replace_time"`
	LotStatus       int32  `json:"lot_status"`
	UserStatus      int32  `json:"user_status"`
	Awards          []struct {
		GiftID   int32  `json:"gift_id"`
		GiftName string `json:"gift_name"`
		Num      int32  `json:"num"`
	} `json:"awards"`
	LotConfigID int32  `json:"lot_config_id"`
	TotalPrice  int32  `json:"total_price"`
	WaitNum     int32  `json:"wait_num"`
	UID         int32  `json:"uid"`
	RoomID      string `json:"roomid"`
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
