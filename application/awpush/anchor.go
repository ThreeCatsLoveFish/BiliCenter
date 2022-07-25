package awpush

import (
	"encoding/json"
	"fmt"
	"net/url"
	"subcenter/domain"
	"subcenter/domain/pull"
	"subcenter/domain/push"
	"subcenter/infra"
	"subcenter/infra/conf"
	"subcenter/infra/dto"
	"subcenter/infra/log"
	"sync/atomic"
	"time"
)

// filterCheckLottery abort blacklist lottery
func filterCheckLottery(anchor dto.AnchorMsg) bool {
	// Need to send gift
	if len(anchor.Data.GiftName) > 0 {
		return true
	}
	// Award is meaningless
	for _, pat := range conf.BiliConf.Filter.WordsPat {
		if pat.MatchString(anchor.Data.AwardName) {
			return true
		}
	}
	// Live room is not safe
	for _, id := range conf.BiliConf.Filter.Rooms {
		if anchor.Data.RoomID == id {
			return true
		}
	}
	// Safe lottery
	return false
}

// joinLottery refers to bilibili live lottery
func joinLottery(client *AWPushClient, anchor dto.AnchorMsg) {
	if filterCheckLottery(anchor) {
		return
	}
	rawUrl := "https://api.live.bilibili.com/xlive/lottery-interface/v1/Anchor/Join"
	data := url.Values{
		"id":       []string{fmt.Sprint(anchor.Data.ID)},
		"platform": []string{"pc"},
	}
	attend := false
	for _, user := range conf.BiliConf.Users {
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
				user.Uid, anchor.Data.ID)
			attend = true
			go func(task domain.Task, timer *time.Timer) {
				<-timer.C
				task.Execute()
			}(domain.Task{
				Pull: pull.NewBiliPull(anchor.Data.RoomID, user.Uid),
				Push: push.NewPush(user.Push),
			}, time.NewTimer(time.Duration(anchor.Data.Time+5)*time.Second))
		} else {
			log.Info("User %d join lottery %d failed because %s",
				user.Uid, anchor.Data.ID, resp.Message)
			if resp.Message == "未登录" && user.Login {
				user.Login = false
				pushEnd := push.NewPush(user.Push)
				pushEnd.Submit(push.Data{
					Title:   "# Cookie失效",
					Content: fmt.Sprintf("用户 %d Cookie失效", user.Uid),
				})
			}
		}
	}
	if attend {
		atomic.AddInt32(&client.join, 1)
	}
}

// HandleAnchorData deal with anchor lottery message
func HandleAnchorData(client *AWPushClient, msg []byte) error {
	var anchor dto.AnchorMsg
	if err := json.Unmarshal(msg, &anchor); err != nil {
		log.Error("Unmarshal AnchorMsg error: %v, raw data: %s", err, string(msg))
		client.sleep.Reset(time.Microsecond)
		return err
	}
	client.sleep.Reset(time.Microsecond)
	atomic.AddInt32(&client.recv, 1)
	go joinLottery(client, anchor)
	return nil
}
