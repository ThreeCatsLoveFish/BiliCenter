package awpush

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"subcenter/application"
	"subcenter/domain/pull"
	"subcenter/domain/push"
	"subcenter/infra"
	"subcenter/infra/dto"
	"time"

	"github.com/gorilla/websocket"
)

type Handler func(*websocket.Conn, []byte, *time.Timer) error

var handlerMap map[string]Handler

func init() {
	registerHandler("HAND_OUT_TASKS", HandleTasks)
	registerHandler("HAND_OUT_ANCHOR_DATA", HandleAnchorData)
}

func registerHandler(name string, handler Handler) {
	if handlerMap == nil {
		handlerMap = make(map[string]Handler)
	}
	handlerMap[name] = handler
}

func getHandler(name string) Handler {
	if handler, ok := handlerMap[name]; ok {
		return handler
	}
	return func(_ *websocket.Conn, _ []byte, timer *time.Timer) error {
		timer.Reset(time.Second)
		return nil
	}
}

// HandleMsg deal with all kinds of message
func HandleMsg(conn *websocket.Conn, timer *time.Timer) error {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Default().Printf("ReadMessage error: %v\n", err)
		return nil
	}
	// Process pong signal
	if string(msg) == "pong" {
		timer.Reset(time.Second)
		return nil
	}
	// Obtain raw data bytes
	raw := infra.PakoInflate(msg)
	var rawMsg dto.RawMsg
	if err = json.Unmarshal(raw, &rawMsg); err != nil {
		log.Default().Printf("Unmarshal error, raw data: %v, error: %v\n", raw, err)
		return err
	}
	if rawMsg.Code != 0 {
		log.Default().Printf("Server code not zero, error exist!")
	}
	// Handle each kind of message
	return getHandler(rawMsg.Type)(conn, raw, timer)
}

// taskCallBack send callback response
func taskCallBack(conn *websocket.Conn, task dto.TaskMsg) error {
	// Sleep before execute
	timer := time.NewTimer(time.Duration(task.Data.SleepTime) * time.Millisecond)
	<-timer.C
	// Send callback message
	resp := dto.Callback{
		Code:   "GET_TASK",
		Uid:    biliConfig.Uid,
		Secret: task.Data.Secret,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		log.Default().Printf("Marshal error: %v\n", err)
		return err
	}
	err = conn.WriteMessage(
		websocket.BinaryMessage,
		infra.PakoDeflate(data),
	)
	if err != nil {
		log.Default().Printf("Callback send error: %v\n", err)
		return err
	}
	return nil
}

// HandleTasks deal with poll task message
func HandleTasks(conn *websocket.Conn, msg []byte, timer *time.Timer) error {
	var task dto.TaskMsg
	if err := json.Unmarshal(msg, &task); err != nil {
		log.Default().Printf("Unmarshal TaskMsg error: %v, raw data: %s\n", err, string(msg))
		return err
	}
	timer.Reset(time.Duration(task.Data.SleepTime) * time.Millisecond)
	go taskCallBack(conn, task)
	return nil
}

// filterCheckLottery abort blacklist lottery
func filterCheckLottery(anchor dto.AnchorMsg) bool {
	// Need to send gift
	if len(anchor.Data.GiftName) > 0 {
		return true
	}
	// Award is meaningless
	for _, pat := range biliConfig.Filter.WordsPat {
		if pat.MatchString(anchor.Data.AwardName) {
			return true
		}
	}
	// Live room is not safe
	for _, id := range biliConfig.Filter.Rooms {
		if anchor.Data.RoomId == id {
			return true
		}
	}
	// Safe lottery
	return false
}

// joinLottery refers to bilibili live lottery
func joinLottery(conn *websocket.Conn, anchor dto.AnchorMsg) {
	if filterCheckLottery(anchor) {
		return
	}
	rawUrl := "https://api.live.bilibili.com/xlive/lottery-interface/v1/Anchor/Join"
	data := url.Values{
		"id":       []string{fmt.Sprint(anchor.Data.Id)},
		"platform": []string{"pc"},
	}
	for _, user := range biliConfig.Users {
		body, err := infra.PostFormWithCookie(rawUrl, user.Cookie, data)
		if err != nil {
			log.Default().Printf("PostFormWithCookie error: %v, raw data: %v\n", err, data)
			continue
		}
		var resp dto.BiliBaseResp
		if err = json.Unmarshal(body, &resp); err != nil {
			log.Default().Printf("Unmarshal BiliJoinResp error: %v, raw data: %v\n", err, body)
		}
		log.Default().Printf("Lottery: %v, Response: %v", anchor, resp)
		task := application.Task{
			Pull: pull.NewBiliPull(anchor.Data.RoomId, user.Uid),
			Push: push.NewPush(user.Push),
		}
		go application.GlobalTaskCenter.AddDelay(
			task, time.Duration(anchor.Data.Time)*time.Second,
		)
	}
}

// HandleAnchorData deal with anchor lottery message
func HandleAnchorData(conn *websocket.Conn, msg []byte, timer *time.Timer) error {
	var anchor dto.AnchorMsg
	if err := json.Unmarshal(msg, &anchor); err != nil {
		log.Default().Printf("Unmarshal AnchorMsg error: %v, raw data: %s\n", err, string(msg))
		return err
	}
	timer.Reset(time.Second)
	go joinLottery(conn, anchor)
	return nil
}
