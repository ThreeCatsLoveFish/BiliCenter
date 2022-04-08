package awpush

import (
	"encoding/json"
	"fmt"
	"net/url"
	"subcenter/infra"
	"subcenter/infra/dto"
	"subcenter/infra/log"
	"time"
)

// joinRedPocket refers to bilibili live lottery
func joinRedPocket(client *AWPushClient, redPocket dto.RedPocketMsg) {
	rawUrl := "https://api.live.bilibili.com/xlive/lottery-interface/v1/popularityRedPocket/RedPocketDraw"
	// FIXME: modify this part to unify type
	var roomId string
	switch val := redPocket.Data.RoomID.(type) {
	case int32, string:
		roomId = fmt.Sprint(val)
	}
	data := url.Values{
		"ruid":       []string{fmt.Sprint(redPocket.Data.UID)},
		"room_id":    []string{roomId},
		"lot_id":     []string{fmt.Sprint(redPocket.Data.LotteryID)},
		"spm_id":     []string{"444.8.red_envelope.extract"},
		"session_id": []string{""},
		"jump_from":  []string{""},
	}
	for _, user := range biliConfig.Users {
		body, err := infra.PostFormWithCookie(rawUrl, user.Cookie, data)
		if err != nil {
			log.Error("PostFormWithCookie error: %v, raw data: %v", err, data)
			continue
		}
		var resp dto.BiliBaseResp
		if err = json.Unmarshal(body, &resp); err != nil {
			log.Error("Unmarshal BiliBaseResp error: %v, raw data: %v", err, body)
		}
		if resp.Code == 0 {
			log.Info("User %d join redpocket %d success",
				user.Uid, redPocket.Data.LotteryID)
		} else {
			log.Info("User %d join redpocket %d failed because %s",
				user.Uid, redPocket.Data.LotteryID, resp.Message)
		}
	}
}

// HandleRedPocket deal with red pocket message
func HandleRedPocket(client *AWPushClient, msg []byte) error {
	log.Debug("Red pocket data is %v", string(msg))
	var redPocket dto.RedPocketMsg
	if err := json.Unmarshal(msg, &redPocket); err != nil {
		log.Error("Unmarshal RedPocketMsg error: %v, raw data: %s", err, string(msg))
		return err
	}
	client.sleep.Reset(time.Microsecond)
	go joinRedPocket(client, redPocket)
	return nil
}
