package dto

// BiliBaseResp is basic response body of all bilibili API
type BiliBaseResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Msg     string `json:"msg,omitempty"`
}

// User info of bilibili
type User struct {
	UID   int    `json:"uid"`
	Uname string `json:"uname"`
	Level int    `json:"level"`
	Color int    `json:"color"`
}

// Anchor represents live anchor lottery
type Anchor struct {
	ID           int    `json:"id"`
	RoomID       int    `json:"room_id"`
	Status       int    `json:"status"`
	AwardName    string `json:"award_name"`
	AwardNum     int    `json:"award_num"`
	AwardImage   string `json:"award_image"`
	Danmu        string `json:"danmu"`
	Time         int    `json:"time"`
	CurrentTime  int    `json:"current_time"`
	JoinType     int    `json:"join_type"`
	RequireType  int    `json:"require_type"`
	RequireValue int    `json:"require_value"`
	RequireText  string `json:"require_text"`
	GiftID       int    `json:"gift_id"`
	GiftName     string `json:"gift_name"`
	GiftNum      int    `json:"gift_num"`
	GiftPrice    int    `json:"gift_price"`
	CurGiftNum   int    `json:"cur_gift_num"`
	GoawayTime   int    `json:"goaway_time"`
	AwardUsers   []User `json:"award_users"`
	ShowPanel    int    `json:"show_panel"`
	LotStatus    int    `json:"lot_status"`
	SendGift     int    `json:"send_gift_ensure"`
	GoodsID      int    `json:"goods_id"`
}

// RedPocket represents red pocket lottery
type RedPocket struct {
	LotteryID       int    `json:"lot_id"`
	SenderUID       int    `json:"sender_uid"`
	SenderName      string `json:"sender_name"`
	JoinRequirement int    `json:"join_requirement"`
	Danmu           string `json:"danmu"`
	CurrentTime     int    `json:"current_time"`
	StartTime       int    `json:"start_time"`
	EndTime         int    `json:"end_time"`
	LastTime        int    `json:"last_time"`
	RemoveTime      int    `json:"remove_time"`
	ReplaceTime     int    `json:"replace_time"`
	LotStatus       int    `json:"lot_status"`
	UserStatus      int    `json:"user_status"`
	Awards          []struct {
		GiftID   int    `json:"gift_id"`
		GiftName string `json:"gift_name"`
		Num      int    `json:"num"`
	} `json:"awards"`
	LotConfigID int         `json:"lot_config_id"`
	TotalPrice  int         `json:"total_price"`
	WaitNum     int         `json:"wait_num"`
	UID         int         `json:"uid"`
	RoomID      interface{} `json:"roomid"`
}

// BiliAnchor is response body of anchor info
type BiliAnchor struct {
	BiliBaseResp
	Data Anchor `json:"data"`
}

type BiliNewTag struct {
	BiliBaseResp
	Data struct {
		TagId int `json:"tagid"`
	} `json:"data"`
}

type BiliListTag struct {
	BiliBaseResp
	Data []struct {
		Name  string `json:"name"`
		TagId int    `json:"tagid"`
		Count int    `json:"count"`
		Tip   string `json:"tip"`
	} `json:"data"`
}

type BiliRelation struct {
	BiliBaseResp
	Data []struct {
		Mid int `json:"mid"`
	} `json:"data"`
}
