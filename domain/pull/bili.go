package pull

import (
	"encoding/json"
	"fmt"
	"net/url"
	"subcenter/domain/push"
	"subcenter/infra"
	"subcenter/infra/dto"
	"subcenter/infra/log"
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
	data, err := infra.Get(rawUrl, "", params)
	if err != nil {
		log.Error("Get error: %v", err)
		return nil, err
	}
	return data, err
}

func (pull BiliPull) Obtain() ([]push.Data, error) {
	var data []push.Data
	body, err := pull.getAwardUser()
	if err != nil {
		log.Error("getAwardUser error: %v", err)
		return nil, err
	}
	var resp dto.BiliAnchor
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliAnchorResp error: %v", err)
		return nil, err
	}
	for _, user := range resp.Data.AwardUsers {
		if user.Uid == pull.uid {
			data = append(data, push.Data{
				Title: "# 天选中奖",
				Content: fmt.Sprintf(
					"用户: %d\n\n房间号: %d\n\n中奖物品: %s",
					user.Uid, resp.Data.RoomId, resp.Data.AwardName,
				),
			})
			log.Info("[LUCK] User %d get award %s",
				user.Uid, resp.Data.AwardName)
		}
	}
	if len(data) == 0 {
		log.Info("Lottery id %d no award", resp.Data.Id)
	}
	return data, nil
}
