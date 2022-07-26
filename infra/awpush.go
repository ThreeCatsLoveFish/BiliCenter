package infra

import (
	"encoding/json"
	"fmt"
	"subcenter/infra/dto"
	"subcenter/infra/log"
	"sync"

	"github.com/gorilla/websocket"
)

var WSMutex sync.Mutex

// Ping send Ping message
func Ping(conn *websocket.Conn) error {
	WSMutex.Lock()
	defer WSMutex.Unlock()
	return conn.WriteMessage(websocket.TextMessage, []byte("ping"))
}

// Pong receive Pong message
func Pong(conn *websocket.Conn) error {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Error("ReadMessage error: %v", err)
		return nil
	}
	if string(msg) == "pong" {
		return nil
	}
	return fmt.Errorf("no pong received")
}

// Verify used for connect awpush
func Verify(conn *websocket.Conn, uid, apiKey string) error {
	data := dto.Verify{
		Code:   "VERIFY_APIKEY",
		Uid:    uid,
		Apikey: apiKey,
	}
	dataStr, err := json.Marshal(data)
	if err != nil {
		log.Error("Marshal error: %v", err)
		return err
	}
	WSMutex.Lock()
	err = conn.WriteMessage(websocket.BinaryMessage, PakoDeflate(dataStr))
	WSMutex.Unlock()
	if err != nil {
		log.Error("Verify error: %v", err)
		return err
	}
	return nil
}
