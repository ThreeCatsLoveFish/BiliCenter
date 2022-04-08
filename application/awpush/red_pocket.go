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
	data := url.Values{
		"ruid": []string{fmt.Sprint(redPocket.Data.RoomUid)}, 
        "room_id": []string{fmt.Sprint(redPocket.Data.RoomId)},
        "lot_id": []string{fmt.Sprint(redPocket.Data.LotteryId)},
        "spm_id": []string{"444.8.red_envelope.extract"},
        "session_id": []string{""},
        "jump_from": []string{""},
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
			log.Info("User %d join lottery %d success",
				user.Uid, redPocket.Data.LotteryId)
		} else {
			log.Info("User %d join lottery %d failed because %s",
				user.Uid, redPocket.Data.LotteryId, resp.Message)
		}
	}
}

// HandleRedPocket deal with red pocket message
func HandleRedPocket(client *AWPushClient, msg []byte) error {
	log.Debug("Red pocket data is %v", msg)
	var redPocket dto.RedPocketMsg
	if err := json.Unmarshal(msg, &redPocket); err != nil {
		log.Error("Unmarshal AnchorMsg error: %v, raw data: %s", err, string(msg))
		return err
	}
	client.sleep.Reset(time.Microsecond)
	go joinRedPocket(client, redPocket)
	return nil
}