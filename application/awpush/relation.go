package awpush

import (
	"encoding/json"
	"fmt"
	"net/url"
	"subcenter/infra"
	"subcenter/infra/conf"
	"subcenter/infra/dto"
	"subcenter/infra/log"
)

const tagName = "BLTH天选关注UP"

// createNewTag open a new tag with prefix of BLTH
func createNewTag(user conf.User) int {
	rawUrl := "https://api.bilibili.com/x/relation/tag/create"
	data := url.Values{
		"tag":  []string{tagName},
		"csrf": []string{user.Csrf},
	}
	body, err := infra.PostFormWithCookie(rawUrl, user.Cookie, data)
	if err != nil {
		log.Error("PostFormWithCookie error: %v, raw data: %v", err, data)
		return 0
	}
	var resp dto.BiliNewTag
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliNewTag error: %v, raw data: %v", err, body)
		return 0
	}
	if resp.Code != 0 {
		log.Error("createNewTag error: %v, raw data: %v", resp.Message, data)
		return 0
	}
	log.Info("create tag for user %d, id is %d", user.Uid, resp.Data.TagId)
	return resp.Data.TagId
}

// getTagId find the tag id of BLTH, if not found then create new tag
func getTagId(user conf.User) int {
	rawUrl := "https://api.bilibili.com/x/relation/tags"
	body, err := infra.Get(rawUrl, user.Cookie, nil)
	if err != nil {
		log.Error("Get error: %v, cookie: %v", err, user.Cookie)
	}
	var resp dto.BiliListTag
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliListTag error: %v, raw data: %v", err, body)
	}
	for _, tag := range resp.Data {
		if tag.Name == tagName {
			return tag.TagId
		}
	}
	// Not found
	return createNewTag(user)
}

// moveUser update user relation
func moveUser(user conf.User, fids string) error {
	rawUrl := "https://api.bilibili.com/x/relation/tags/moveUsers"
	data := url.Values{
		"beforeTagids": []string{"0"},
		"afterTagids":  []string{fmt.Sprint(getTagId(user))},
		"fids":         []string{fids},
		"csrf":         []string{user.Csrf},
	}
	body, err := infra.PostFormWithCookie(rawUrl, user.Cookie, data)
	if err != nil {
		log.Error("PostFormWithCookie error: %v, raw data: %v", err, data)
		return err
	}
	var resp dto.BiliBaseResp
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliBaseResp error: %v, raw data: %v", err, body)
		return err
	}
	if resp.Code != 0 {
		msg := fmt.Sprintf("moveUsers error: %v, raw data: %v", resp.Message, data)
		log.Error(msg)
		return fmt.Errorf(msg)
	}
	return nil
}

// getRelation obtains users in default tag list and move to BLTH group
func getRelation(user conf.User) (string, error) {
	rawUrl := "https://api.bilibili.com/x/relation/tag"
	data := url.Values{
		"mid":   []string{fmt.Sprint(user.Uid)},
		"tagid": []string{"0"},
		"pn":    []string{"1"},
		"ps":    []string{"20"},
	}
	body, err := infra.Get(rawUrl, user.Cookie, data)
	if err != nil {
		log.Error("Get error: %v, raw data: %v", err, data)
		return "", err
	}
	var resp dto.BiliRelation
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Error("Unmarshal BiliRelation error: %v, raw data: %v", err, body)
		return "", err
	}
	var fids string
	for idx, mid := range resp.Data {
		if idx == 0 {
			fids += fmt.Sprintf("%d", mid.Mid)
		} else {
			fids += fmt.Sprintf(",%d", mid.Mid)
		}
	}
	return fids, nil
}

// UpdateRelation traverse all account and update relation
func UpdateRelation() interface{} {
	type result struct {
		Uid   int   `json:"uid"`
		Error error `json:"err"`
	}
	fail := make([]result, 0)
	for _, user := range conf.BiliConf.Users {
		fids, err := getRelation(user)
		if err != nil {
			data := result{
				Uid:   user.Uid,
				Error: err,
			}
			log.Error("getRelation failed, detail: %v", data)
			fail = append(fail, data)
			continue
		}
		if len(fids) == 0 {
			log.Info("No relation need update for user %d", user.Uid)
			continue
		}
		if err = moveUser(user, fids); err != nil {
			data := result{
				Uid:   user.Uid,
				Error: err,
			}
			log.Error("moveUser failed, detail: %v", data)
			fail = append(fail, data)
			continue
		}
		log.Info("Update relation success for user %d", user.Uid)
	}
	if len(fail) == 0 {
		return map[string]interface{}{
			"code":   0,
			"status": "Update relation success",
		}
	}
	return map[string]interface{}{
		"code":   5,
		"status": "Update relation fail",
		"data":   fail,
	}
}
