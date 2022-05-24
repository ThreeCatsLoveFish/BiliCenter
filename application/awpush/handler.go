package awpush

import (
	"encoding/json"
	"subcenter/infra"
	"subcenter/infra/dto"
	"subcenter/infra/log"
	"time"

	"github.com/gorilla/websocket"
)

type Handler func(*AWPushClient, []byte) error

var handlerMap map[string]Handler

func init() {
	registerHandler("HAND_OUT_TASKS", HandleTasks)
	registerHandler("HAND_OUT_ANCHOR_DATA", HandleAnchorData)
	registerHandler("HAND_OUT_POPULARITY_REDPOCKET_DATA", HandleRedPocket)
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
	return func(client *AWPushClient, _ []byte) error {
		client.sleep.Reset(time.Microsecond)
		return nil
	}
}

// HandleMsg deal with all kinds of message
func HandleMsg(client *AWPushClient) error {
	_, msg, err := client.conn.ReadMessage()
	if err != nil {
		log.Error("ReadMessage error: %v", err)
		return err
	}
	// Process pong signal
	if string(msg) == "pong" {
		log.Debug("Heartbeat received")
		client.sleep.Reset(time.Microsecond)
		return nil
	}
	// Obtain raw data bytes
	raw := infra.PakoInflate(msg)
	var rawMsg dto.RawMsg
	if err = json.Unmarshal(raw, &rawMsg); err != nil {
		log.Error("Unmarshal error, raw data: %v, error: %v", raw, err)
		return err
	}
	if rawMsg.Code != 0 {
		log.Error("Server code not zero, error exist!")
	}
	log.Debug("Get message %v", rawMsg)
	// Handle each kind of message
	return getHandler(rawMsg.Type)(client, raw)
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
		log.Error("Marshal error: %v", err)
		return err
	}
	err = conn.WriteMessage(websocket.BinaryMessage, infra.PakoDeflate(data))
	if err != nil {
		log.Error("Callback send error: %v", err)
		return err
	}
	return nil
}

// HandleTasks deal with poll task message
func HandleTasks(client *AWPushClient, msg []byte) error {
	var task dto.TaskMsg
	if err := json.Unmarshal(msg, &task); err != nil {
		log.Error("Unmarshal TaskMsg error: %v, raw data: %s", err, string(msg))
		return err
	}
	client.sleep.Reset(15 * time.Second)
	go taskCallBack(client.conn, task)
	return nil
}
