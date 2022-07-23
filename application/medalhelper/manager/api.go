package manager

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"subcenter/infra"
	"subcenter/infra/conf"
	"subcenter/infra/dto"
	"subcenter/infra/log"
)

func LoginVerify(user conf.User) (dto.BiliAccountResp, error) {
	rawUrl := "https://app.bilibili.com/x/v2/account/mine"
	data := url.Values{
		"platform":   []string{"pc"},
		"csrf":       []string{user.Csrf},
		"csrf_token": []string{user.Csrf},
	}
	var resp dto.BiliAccountResp
	body, err := infra.Get(rawUrl, user.Cookie, data)
	if err != nil {
		log.Error("LoginVerify error: %v, data: %v", err, data)
		return resp, err
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliAccountResp error: %v, raw data: %v", err, body)
		return resp, err
	}
	return resp, nil
}

func SignIn(user conf.User) (string, error) {
	rawUrl := "https://api.live.bilibili.com/rc/v1/Sign/doSign"
	data := url.Values{
		"platform":   []string{"pc"},
		"csrf":       []string{user.Csrf},
		"csrf_token": []string{user.Csrf},
	}
	body, err := infra.Get(rawUrl, user.Cookie, data)
	if err != nil {
		log.Error("SignIn error: %v, data: %v", err, data)
		return "", err
	}
	return string(body), nil
}

func GetUserInfo(user conf.User) (dto.BiliLiveUserInfo, error) {
	rawUrl := "https://api.live.bilibili.com/xlive/app-ucenter/v1/user/get_user_info"
	data := url.Values{
		"platform":   []string{"pc"},
		"csrf":       []string{user.Csrf},
		"csrf_token": []string{user.Csrf},
	}
	body, err := infra.Get(rawUrl, user.Cookie, data)
	var resp dto.BiliLiveUserInfo
	if err != nil {
		log.Error("GetUserInfo error: %v, data: %v", err, data)
		return resp, err
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliLiveUserInfo error: %v, raw data: %v", err, body)
		return resp, err
	}
	return resp, nil
}

func GetMedal(user conf.User) ([]dto.MedalInfo, bool) {
	medals := make([]dto.MedalInfo, 0, 20)
	wear := false
	page := 1
	for {
		rawUrl := "https://api.live.bilibili.com/xlive/app-ucenter/v1/fansMedal/panel"
		data := url.Values{
			"platform":   []string{"pc"},
			"csrf":       []string{user.Csrf},
			"csrf_token": []string{user.Csrf},
			"page":       []string{fmt.Sprint(page)},
			"page_size":  []string{"50"},
		}
		body, err := infra.PostFormWithCookie(rawUrl, user.Cookie, data)
		if err != nil {
			log.Error("GetFansMedalAndRoomID error: %v, data: %v", err, data)
			return medals, wear
		}
		var resp dto.BiliMedalResp
		if err = json.Unmarshal(body, &resp); err != nil {
			log.Error("Unmarshal BiliMedalResp error: %v, raw data: %v", err, body)
			return medals, wear
		}
		if len(resp.Data.SpecialList) > 0 {
			wear = true
		}
		medals = append(medals, resp.Data.SpecialList...)
		medals = append(medals, resp.Data.List...)
		if len(resp.Data.List) == 0 {
			break
		}
		page++
	}
	return medals, wear
}

func WearMedal(user conf.User, medalId int) bool {
	rawUrl := "https://api.live.bilibili.com/xlive/app-ucenter/v1/fansMedal/wear"
	data := url.Values{
		"platform":   []string{"pc"},
		"medal_id":   []string{fmt.Sprint(medalId)},
		"csrf":       []string{user.Csrf},
		"csrf_token": []string{user.Csrf},
	}
	body, err := infra.PostFormWithCookie(rawUrl, user.Cookie, data)
	if err != nil {
		log.Error("WearMedal error: %v, data: %v", err, data)
	}
	var resp dto.BiliBaseResp
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliBaseResp error: %v, raw data: %v", err, body)
	}
	return resp.Code == 0
}

func TakeoffMedal(user conf.User) bool {
	rawUrl := "https://api.live.bilibili.com/xlive/app-ucenter/v1/fansMedal/take_off"
	data := url.Values{
		"platform":   []string{"pc"},
		"csrf":       []string{user.Csrf},
		"csrf_token": []string{user.Csrf},
	}
	body, err := infra.PostFormWithCookie(rawUrl, user.Cookie, data)
	if err != nil {
		log.Error("TakeoffMedal error: %v, data: %v", err, data)
	}
	var resp dto.BiliBaseResp
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliBaseResp error: %v, raw data: %v", err, body)
	}
	return resp.Code == 0
}

func LikeInteract(user conf.User, roomId int) bool {
	rawUrl := "https://api.live.bilibili.com/xlive/web-ucenter/v1/interact/likeInteract"
	data := url.Values{
		"platform":   []string{"pc"},
		"roomid":     []string{fmt.Sprint(roomId)},
		"csrf":       []string{user.Csrf},
		"csrf_token": []string{user.Csrf},
	}
	body, err := infra.PostFormWithCookie(rawUrl, user.Cookie, data)
	if err != nil {
		log.Error("LikeInteract error: %v, data: %v", err, data)
	}
	var resp dto.BiliBaseResp
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliBaseResp error: %v, raw data: %v", err, body)
	}
	return resp.Code == 0
}

func SendDanmaku(user conf.User, roomId int) bool {
	rawUrl := "https://api.live.bilibili.com/xlive/app-room/v1/dM/sendmsg"
	data := url.Values{
		"platform":   []string{"pc"},
		"cid":        []string{fmt.Sprint(roomId)},
		"msg":        []string{"牛呀牛呀"},
		"rnd":        []string{GetTimestamp()},
		"color":      []string{"16777215"},
		"fontsize":   []string{"25"},
		"csrf":       []string{user.Csrf},
		"csrf_token": []string{user.Csrf},
	}
	body, err := infra.PostFormWithCookie(rawUrl, user.Cookie, data)
	if err != nil {
		log.Error("GetFansMedalAndRoomID error: %v, data: %v", err, data)
	}
	var resp dto.BiliBaseResp
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliBaseResp error: %v, raw data: %v", err, body)
	}
	return resp.Code == 0
}

func Heartbeat(user conf.User, uuids []string, roomId, upId int) bool {
	rawUrl := "https://live-trace.bilibili.com/xlive/data-interface/v1/heartbeat/mobileHeartBeat"
	data := url.Values{
		"platform":         []string{"pc"},
		"uuid":             []string{uuids[0]},
		"buvid":            []string{strings.ToUpper(RandomString(37))},
		"seq_id":           []string{"1"},
		"room_id":          []string{fmt.Sprint(roomId)},
		"parent_id":        []string{"6"},
		"area_id":          []string{"283"},
		"timestamp":        []string{fmt.Sprintf("%d", time.Now().Unix()-60)},
		"secret_key":       []string{"axoaadsffcazxksectbbb"},
		"watch_time":       []string{"60"},
		"up_id":            []string{fmt.Sprint(upId)},
		"up_level":         []string{"40"},
		"jump_from":        []string{"30000"},
		"gu_id":            []string{strings.ToUpper(RandomString(43))},
		"play_type":        []string{"0"},
		"play_url":         []string{""},
		"s_time":           []string{"0"},
		"data_behavior_id": []string{""},
		"data_source_id":   []string{""},
		"up_session":       []string{fmt.Sprintf("l:one:live:record:%d:%d", roomId, time.Now().Unix()-88888)},
		"visit_id":         []string{strings.ToUpper(RandomString(32))},
		"watch_status":     []string{"%7B%22pk_id%22%3A0%2C%22screen_status%22%3A1%7D"},
		"click_id":         []string{uuids[1]},
		"session_id":       []string{""},
		"player_type":      []string{"0"},
		"client_ts":        []string{GetTimestamp()},
		"csrf":             []string{user.Csrf},
		"csrf_token":       []string{user.Csrf},
	}
	body, err := infra.PostFormWithCookie(rawUrl, user.Cookie, data)
	if err != nil {
		log.Error("Heartbeat error: %v, data: %v", err, data)
	}
	var resp dto.BiliBaseResp
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliBaseResp error: %v, raw data: %v", err, body)
	}
	return resp.Code == 0
}
