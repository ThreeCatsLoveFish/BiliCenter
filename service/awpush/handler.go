package awpush

import (
	"encoding/json"
	"fmt"
	"net/url"
	"subcenter/manager"
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

func HandleMsg(conn *websocket.Conn, timer *time.Timer) error {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		fmt.Printf("ReadMessage error: %v\n", err)
		return nil
	}
	// Process pong signal
	if string(msg) == "pong" {
		timer.Reset(time.Second)
		return nil
	}
	// Obtain raw data bytes
	raw := manager.PakoInflate(msg)
	var rawMsg RawMsg
	if err = json.Unmarshal(raw, &rawMsg); err != nil {
		fmt.Printf("Unmarshal error, raw data: %v, error: %v\n", raw, err)
		return err
	}
	if rawMsg.Code != 0 {
		fmt.Printf("Server code not zero, error exist!")
	}
	// Handle each kind of message
	return getHandler(rawMsg.Type)(conn, raw, timer)
}

// taskCallBack send callback response
func taskCallBack(conn *websocket.Conn, task TaskMsg) error {
	// Sleep before execute
	timer := time.NewTimer(time.Duration(task.Data.SleepTime) * time.Millisecond)
	<-timer.C
	// Send callback message
	resp := Callback{
		Code:   "GET_TASK",
		Uid:    biliConfig.Uid,
		Secret: task.Data.Secret,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Marshal error: %v\n", err)
		return err
	}
	err = conn.WriteMessage(
		websocket.BinaryMessage,
		manager.PakoDeflate(data),
	)
	if err != nil {
		fmt.Printf("Callback send error: %v\n", err)
		return err
	}
	return nil
}

// HandleTasks deal with poll task message
func HandleTasks(conn *websocket.Conn, msg []byte, timer *time.Timer) error {
	var task TaskMsg
	if err := json.Unmarshal(msg, &task); err != nil {
		fmt.Printf("Unmarshal TaskMsg error: %v, raw data: %s\n", err, string(msg))
		return err
	}
	timer.Reset(time.Duration(task.Data.SleepTime) * time.Millisecond)
	go taskCallBack(conn, task)
	return nil
}

// filterCheckLottery abort blacklist lottery
func filterCheckLottery(anchor AnchorMsg) bool {
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

// biliJoinLottery join bilibili live lottery
func biliJoinLottery(conn *websocket.Conn, anchor AnchorMsg) {
	if filterCheckLottery(anchor) {
		return
	}
	rawUrl := "https://api.live.bilibili.com/xlive/lottery-interface/v1/Anchor/Join"
	data := url.Values{
		"id":       []string{fmt.Sprint(anchor.Data.Id)},
		"platform": []string{"pc"},
	}
	for _, cookie := range biliConfig.Cookies {
		body, err := manager.PostFormWithCookie(rawUrl, cookie, data)
		if err != nil {
			fmt.Printf("PostFormWithCookie error: %v, raw data: %v\n", err, data)
			continue
		}
		var resp BiliJoinResp
		if err = json.Unmarshal(body, &resp); err != nil {
			fmt.Printf("Unmarshal BiliJoinResp error: %v, raw data: %v\n", err, body)
		}
		fmt.Printf("Lottery: %v, Response: %v", anchor, resp)
	}
}

// HandleAnchorData deal with anchor lottery message
func HandleAnchorData(conn *websocket.Conn, msg []byte, timer *time.Timer) error {
	var anchor AnchorMsg
	if err := json.Unmarshal(msg, &anchor); err != nil {
		fmt.Printf("Unmarshal AnchorMsg error: %v, raw data: %s\n", err, string(msg))
		return err
	}
	timer.Reset(time.Second)
	go biliJoinLottery(conn, anchor)
	return nil
}
