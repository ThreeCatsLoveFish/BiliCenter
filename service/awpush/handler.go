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
	registerHandler("HAND_OUT_TASKS", handleTasks)
	registerHandler("HAND_OUT_ANCHOR_DATA", handleAnchorData)
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

func handleMsg(conn *websocket.Conn, timer *time.Timer) error {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		fmt.Printf("ReadMessage error: %v\n", err)
		return nil
	}
	// Process pong signal
	if string(msg) == "pong" {
		fmt.Printf("Pong received, reset status\n")
		timer.Reset(time.Second)
		return nil
	}
	// Obtain raw data bytes
	raw := manager.PakoInflate(msg)
	var rawMsg RawMsg
	fmt.Printf("Receive msg: %s\n", string(raw))
	err = json.Unmarshal(raw, &rawMsg)
	if err != nil {
		fmt.Printf("Unmarshal error, raw data: %v, error: %v\n", raw, err)
		return err
	}
	if rawMsg.Code != 0 {
		fmt.Printf("Server code not zero, error exist!")
	}
	// Handle each kind of message
	return getHandler(rawMsg.Type)(conn, raw, timer)
}

// handleTasks deal with poll task message
func handleTasks(conn *websocket.Conn, msg []byte, timer *time.Timer) error {
	var task TaskMsg
	if err := json.Unmarshal(msg, &task); err != nil {
		fmt.Printf("Unmarshal TaskMsg error: %v, raw data: %s\n", err, string(msg))
		return err
	}
	timer.Reset(time.Duration(task.Data.SleepTime) * time.Millisecond)
	go taskCallBack(conn, task)
	return nil
}

// taskCallBack send callback response
func taskCallBack(conn *websocket.Conn, task TaskMsg) error {
	// Sleep before execute
	timer := time.NewTimer(time.Duration(task.Data.SleepTime) * time.Millisecond)
	<-timer.C
	// Send callback message
	resp := Resp{
		Code:   "GET_TASK",
		Uid:    "12309253",
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

// handleAnchorData deal with anchor lottery message
func handleAnchorData(conn *websocket.Conn, msg []byte, timer *time.Timer) error {
	var anchor AnchorMsg
	err := json.Unmarshal(msg, &anchor)
	if err != nil {
		fmt.Printf("Unmarshal AnchorMsg error: %v, raw data: %s\n", err, string(msg))
		return err
	}
	timer.Reset(time.Second)
	go joinLottery(conn, anchor)
	return nil
}

// taskCallBack send callback response
func joinLottery(conn *websocket.Conn, anchor AnchorMsg) error {
	rawUrl := "https://api.live.bilibili.com/xlive/lottery-interface/v1/Anchor/Join"
	data := url.Values{
		"id":         []string{fmt.Sprint(anchor.Data.Id)},
		"platform":   []string{"pc"},
		"csrf":       []string{"973fa29433bb4741d8d89c7cc5e93f22"},
		"csrf_token": []string{"973fa29433bb4741d8d89c7cc5e93f22"},
	}
	cookie := "_uuid=46371455-3BAB-B9C5-15E5-B2C93B4F9D7C73100infoc; buvid3=9710689A-A75E-4705-99F0-EEB514352484148822infoc; fingerprint=e580f0a4e0630916ac9a68f003b586f2; buvid_fp=cbbc1e3fd2572c2bbc3749a5110e72d1; buvid_fp_plain=9710689A-A75E-4705-99F0-EEB514352484148822infoc; bp_video_offset_12309253=626586469511221100; CURRENT_FNVAL=2000; blackside_state=1; rpdid=|(k|kYYYukYR0J'uYk|uRk))l; PVID=4; LIVE_BUVID=AUTO9316212512347742; fingerprint3=d3b52b79169367f8d9177411386219bb; fingerprint_s=3857c3522fc3c3f2f1c68c81e5a07930; CURRENT_QUALITY=112; CURRENT_BLACKGAP=1; video_page_version=v_old_home; sid=6de9gsar; SESSDATA=645671bd%2C1651557113%2C329dd%2Ab1; bili_jct=973fa29433bb4741d8d89c7cc5e93f22; DedeUserID=12309253; DedeUserID__ckMd5=665a54d6c2ac1031; i-wanna-go-back=2; b_ut=5; buvid4=5EFCA494-C578-942F-7971-5F7636B3D46331659-022012821-DuABVwrxQm8XnuNNhp2vNw%3D%3D; bp_t_offset_12309253=626395399302499850; innersign=1; b_lsid=104F586BF_17EF3A185F1; bsource=search_bing"
	return manager.PostFormWithCookie(rawUrl, data, cookie)
}
