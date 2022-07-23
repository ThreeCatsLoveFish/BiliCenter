package dto

// BiliDataResp only check status and data
type BiliDataResp struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	TTL     int         `json:"ttl,omitempty"`
	Msg     string      `json:"msg,omitempty"`
	Data    interface{} `json:"data"`
}

type MedalInfo struct {
	Medal struct {
		UID              int    `json:"uid"`
		TargetID         int    `json:"target_id"`
		TargetName       string `json:"target_name"`
		MedalID          int    `json:"medal_id"`
		Level            int    `json:"level"`
		MedalName        string `json:"medal_name"`
		MedalColor       int    `json:"medal_color"`
		Intimacy         int    `json:"intimacy"`
		NextIntimacy     int    `json:"next_intimacy"`
		DayLimit         int    `json:"day_limit"`
		TodayFeed        int    `json:"today_feed"`
		MedalColorStart  int    `json:"medal_color_start"`
		MedalColorEnd    int    `json:"medal_color_end"`
		MedalColorBorder int    `json:"medal_color_border"`
		IsLighted        int    `json:"is_lighted"`
		GuardLevel       int    `json:"guard_level"`
		WearingStatus    int    `json:"wearing_status"`
		MedalIconID      int    `json:"medal_icon_id"`
		MedalIconURL     string `json:"medal_icon_url"`
		GuardIcon        string `json:"guard_icon"`
		HonorIcon        string `json:"honor_icon"`
		CanDelete        bool   `json:"can_delete"`
	} `json:"medal"`
	AnchorInfo struct {
		NickName string `json:"nick_name"`
		Avatar   string `json:"avatar"`
		Verify   int    `json:"verify"`
	} `json:"anchor_info"`
	Superscript interface{} `json:"superscript"`
	RoomInfo    struct {
		RoomID       int    `json:"room_id"`
		LivingStatus int    `json:"living_status"`
		URL          string `json:"url"`
	} `json:"room_info"`
}

var DefaultMedal MedalInfo

// BiliMedalResp obtain the response with all medal info
type BiliMedalResp struct {
	BiliBaseResp
	Data struct {
		List        []MedalInfo `json:"list"`
		SpecialList []MedalInfo `json:"special_list"`
		BottomBar   interface{} `json:"bottom_bar"`
		PageInfo    struct {
			Number          int  `json:"number"`
			CurrentPage     int  `json:"current_page"`
			HasMore         bool `json:"has_more"`
			NextPage        int  `json:"next_page"`
			NextLightStatus int  `json:"next_light_status"`
			TotalPage       int  `json:"total_page"`
		} `json:"page_info"`
		TotalNumber int `json:"total_number"`
		HasMedal    int `json:"has_medal"`
	} `json:"data"`
}

// BiliUserInfo represent user live info
type BiliLiveUserInfo struct {
	BiliBaseResp
	Data struct {
		UID              int     `json:"uid"`
		Uname            string  `json:"uname"`
		Face             string  `json:"face"`
		BillCoin         float64 `json:"billCoin"`
		Silver           int     `json:"silver"`
		Gold             int     `json:"gold"`
		Achieve          int     `json:"achieve"`
		Vip              int     `json:"vip"`
		Svip             int     `json:"svip"`
		UserLevel        int     `json:"user_level"`
		UserNextLevel    int     `json:"user_next_level"`
		UserIntimacy     int     `json:"user_intimacy"`
		UserNextIntimacy int     `json:"user_next_intimacy"`
		IsLevelTop       int     `json:"is_level_top"`
		UserLevelRank    string  `json:"user_level_rank"`
		UserCharged      int     `json:"user_charged"`
		Identification   int     `json:"identification"`
	} `json:"data"`
}

type BiliLiveRoomInfo struct {
	BiliBaseResp
	Data    struct {
		UID              int      `json:"uid"`
		RoomID           int      `json:"room_id"`
		ShortID          int      `json:"short_id"`
		Attention        int      `json:"attention"`
		Online           int      `json:"online"`
		IsPortrait       bool     `json:"is_portrait"`
		Description      string   `json:"description"`
		LiveStatus       int      `json:"live_status"`
		AreaID           int      `json:"area_id"`
		ParentAreaID     int      `json:"parent_area_id"`
		ParentAreaName   string   `json:"parent_area_name"`
		OldAreaID        int      `json:"old_area_id"`
		Background       string   `json:"background"`
		Title            string   `json:"title"`
		UserCover        string   `json:"user_cover"`
		Keyframe         string   `json:"keyframe"`
		IsStrictRoom     bool     `json:"is_strict_room"`
		LiveTime         string   `json:"live_time"`
		Tags             string   `json:"tags"`
		IsAnchor         int      `json:"is_anchor"`
		RoomSilentType   string   `json:"room_silent_type"`
		RoomSilentLevel  int      `json:"room_silent_level"`
		RoomSilentSecond int      `json:"room_silent_second"`
		AreaName         string   `json:"area_name"`
		Pendants         string   `json:"pendants"`
		AreaPendants     string   `json:"area_pendants"`
		HotWords         []string `json:"hot_words"`
		HotWordsStatus   int      `json:"hot_words_status"`
		Verify           string   `json:"verify"`
		NewPendants      struct {
			Frame struct {
				Name       string `json:"name"`
				Value      string `json:"value"`
				Position   int    `json:"position"`
				Desc       string `json:"desc"`
				Area       int    `json:"area"`
				AreaOld    int    `json:"area_old"`
				BgColor    string `json:"bg_color"`
				BgPic      string `json:"bg_pic"`
				UseOldArea bool   `json:"use_old_area"`
			} `json:"frame"`
			Badge struct {
				Name     string `json:"name"`
				Position int    `json:"position"`
				Value    string `json:"value"`
				Desc     string `json:"desc"`
			} `json:"badge"`
			MobileFrame struct {
				Name       string `json:"name"`
				Value      string `json:"value"`
				Position   int    `json:"position"`
				Desc       string `json:"desc"`
				Area       int    `json:"area"`
				AreaOld    int    `json:"area_old"`
				BgColor    string `json:"bg_color"`
				BgPic      string `json:"bg_pic"`
				UseOldArea bool   `json:"use_old_area"`
			} `json:"mobile_frame"`
			MobileBadge interface{} `json:"mobile_badge"`
		} `json:"new_pendants"`
		UpSession            string `json:"up_session"`
		PkStatus             int    `json:"pk_status"`
		PkID                 int    `json:"pk_id"`
		BattleID             int    `json:"battle_id"`
		AllowChangeAreaTime  int    `json:"allow_change_area_time"`
		AllowUploadCoverTime int    `json:"allow_upload_cover_time"`
		StudioInfo           struct {
			Status     int           `json:"status"`
			MasterList []interface{} `json:"master_list"`
		} `json:"studio_info"`
	} `json:"data"`
}

type BiliHeartBeatResp struct {
	BiliBaseResp
	Data struct {
		Timestamp         int    `json:"timestamp"`
		HeartbeatInterval int    `json:"heartbeat_interval"`
		SecretKey         string `json:"secret_key"`
		SecretRule        []int  `json:"secret_rule"`
		PatchStatus       int    `json:"patch_status"`
	} `json:"data"`
}
