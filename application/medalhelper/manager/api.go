package manager

import (
	"encoding/json"
	"fmt"
	"net/url"

	"subcenter/infra"
	"subcenter/infra/conf"
	"subcenter/infra/dto"
	"subcenter/infra/log"
)

func SignIn(user conf.User) (dto.BiliBaseResp, error) {
	rawUrl := "https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/DoSign"
	data := url.Values{
		"platform": []string{"pc"},
	}
	body, err := infra.Get(rawUrl, user.Cookie, data)
	var resp dto.BiliBaseResp
	if err != nil {
		log.Error("GetUserInfo error: %v, data: %v", err, data)
		return resp, err
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliBaseResp error: %v, raw data: %v", err, body)
		return resp, err
	}
	return resp, nil
}

func GetUserInfo(user conf.User) (dto.BiliLiveUserInfo, error) {
	rawUrl := "https://api.live.bilibili.com/xlive/web-ucenter/user/get_user_info"
	data := url.Values{
		"platform": []string{"pc"},
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

func GetRoomInfo(user conf.User, roomId int) (dto.BiliLiveRoomInfo, error) {
	rawUrl := "https://api.live.bilibili.com/room/v1/Room/get_info"
	data := url.Values{
		"room_id": []string{fmt.Sprint(roomId)},
		"from":    []string{"room"},
	}
	body, err := infra.Get(rawUrl, user.Cookie, data)
	var resp dto.BiliLiveRoomInfo
	if err != nil {
		log.Error("GetUserInfo error: %v, data: %v", err, data)
		return resp, err
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliLiveRoomInfo error: %v, raw data: %v", err, body)
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
			"page":      []string{fmt.Sprint(page)},
			"page_size": []string{"50"},
		}
		body, err := infra.Get(rawUrl, user.Cookie, data)
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

func E(user conf.User, uuids []string, roomId int) (dto.BiliLiveRoomInfo, dto.BiliHeartBeatResp) {
	room, _ := GetRoomInfo(user, roomId)
	rawUrl := "https://live-trace.bilibili.com/xlive/data-interface/v1/x25Kn/E"
	data := url.Values{
		"id": []string{fmt.Sprintf("[%d,%d,0,%d]",
			room.Data.ParentAreaID, room.Data.AreaID, room.Data.RoomID)},
		"device":     []string{fmt.Sprintf("[\"%s\",\"%s\"]", user.Buvid, uuids[0])},
		"ts":         []string{GetTimestamp()},
		"is_patch":   []string{"0"},
		"heart_beat": []string{"[]"},
		"ua":         []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0"},
		"csrf":       []string{user.Csrf},
		"csrf_token": []string{user.Csrf},
		"visit_id":   []string{""},
	}
	body, err := infra.PostFormWithCookie(rawUrl, user.Cookie, data)
	if err != nil {
		log.Error("Heartbeat error: %v, data: %v", err, data)
	}
	var resp dto.BiliHeartBeatResp
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliHeartBeatResp error: %v, raw data: %v", err, body)
	}
	return room, resp
}

func X(user conf.User, uuids []string, seq int, room dto.BiliLiveRoomInfo, hb dto.BiliHeartBeatResp) dto.BiliHeartBeatResp {
	rawUrl := "https://live-trace.bilibili.com/xlive/data-interface/v1/x25Kn/X"
	data := url.Values{
		"id": []string{fmt.Sprintf("[%d,%d,%d,%d]",
			room.Data.ParentAreaID, room.Data.AreaID, seq, room.Data.RoomID)},
		"device":     []string{fmt.Sprintf("[\"%s\",\"%s\"]", user.Buvid, uuids[0])},
		"platform":   []string{"web"},
		"ets":        []string{fmt.Sprint(hb.Data.Timestamp)},
		"benchmark":  []string{hb.Data.SecretKey},
		"time":       []string{fmt.Sprint(hb.Data.HeartbeatInterval)},
		"ts":         []string{GetTimestamp()},
		"ua":         []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0"},
		"csrf":       []string{user.Csrf},
		"csrf_token": []string{user.Csrf},
		"visit_id":   []string{""},
	}
	dataStr := fmt.Sprintf(`{"platform":"web","parent_id":%d,"area_id":%d,"seq_id":%d,"room_id":%d,"buvid":"%s","uuid":"%s","ets":%d,"time":%d,"ts":%s}`,
		room.Data.ParentAreaID,
		room.Data.AreaID,
		seq,
		room.Data.RoomID,
		user.Buvid,
		uuids[0],
		hb.Data.Timestamp,
		hb.Data.HeartbeatInterval,
		data["ts"][0],
	)
	sum := CryptoSign(dataStr, hb.Data.SecretKey, hb.Data.SecretRule)
	data["s"] = []string{sum}
	body, err := infra.PostFormWithCookie(rawUrl, user.Cookie, data)
	if err != nil {
		log.Error("Heartbeat error: %v, data: %v", err, data)
	}
	var resp dto.BiliHeartBeatResp
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliHeartBeatResp error: %v, raw data: %v", err, body)
	}
	return resp
}
