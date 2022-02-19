package pull

import (
	"encoding/json"
	"fmt"
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
		"roomid": []string{fmt.Sprint(pull)},
	}
	data, err := infra.GetWithParams(rawUrl, params)
	if err != nil {
		// FIXME: add log here
		return nil, err
	}
	return data, err
}

func (pull BiliPull) Obtain() ([]push.Data, error) {
	var data []push.Data
	body, err := pull.getAwardUser()
	if err != nil {
		// FIXME: add log here
		return nil, err
	}
	var resp dto.BiliAnchorResp
	if err = json.Unmarshal(body, &resp); err != nil {
		// FIXME: add log here
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
		}
	}
	return data, nil
}
