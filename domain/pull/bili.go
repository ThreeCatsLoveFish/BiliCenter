package pull

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"subcenter/domain/push"
	"subcenter/infra"
	"subcenter/infra/dto"
)

type BiliPull struct {
	roomid int32
	uid    int32
}

func NewBiliPull(roomid int32, uid int32) BiliPull {
	return BiliPull{roomid, uid}
}

func (pull BiliPull) getAwardUser() ([]byte, error) {
	rawUrl := "https://api.live.bilibili.com/xlive/lottery-interface/v1/Anchor/Check"
	params := url.Values{
		"roomid": []string{fmt.Sprint(pull.roomid)},
	}
	data, err := infra.GetWithParams(rawUrl, params)
	if err != nil {
		log.Default().Printf("GetWithParams error: %v", err)
		return nil, err
	}
	return data, err
}

func (pull BiliPull) Obtain() ([]push.Data, error) {
	var data []push.Data
	body, err := pull.getAwardUser()
	if err != nil {
		log.Default().Printf("getAwardUser error: %v", err)
		return nil, err
	}
	var resp dto.BiliAnchorResp
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Default().Printf("Unmarshal BiliAnchorResp error: %v", err)
		return nil, err
	}
	for _, user := range resp.Data.AwardUsers {
		if user.Uid == pull.uid {
			data = append(data, push.Data{
				Title: "天选中奖",
				Content: fmt.Sprintf(
					"房间号: %d\n\n中奖物品: %s",
					resp.Data.RoomId, resp.Data.AwardName,
				),
			})
			log.Default().Printf("[LUCK] User %d get award %s",
				user.Uid, resp.Data.AwardName)
		}
	}
	if len(data) == 0 {
		log.Default().Printf("[INFO] Lottery id %d no award", resp.Data.Id)
	}
	return data, nil
}
