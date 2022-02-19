package infra

import (
	"encoding/json"
	"fmt"
	"subcenter/infra/dto"

	"github.com/gorilla/websocket"
)

// Ping send Ping message
func Ping(conn *websocket.Conn) error {
	err := conn.WriteMessage(
		websocket.TextMessage,
		[]byte("ping"),
	)
	if err != nil {
		fmt.Printf("Ping error: %v\n", err)
	}
	return err
}

// Pong receive Pong message
func Pong(conn *websocket.Conn) error {
	_, msg, err := conn.ReadMessage()
	if err != nil {
		fmt.Printf("ReadMessage error: %v\n", err)
		return nil
	}
	if string(msg) == "pong" {
		fmt.Printf("Pong received, reset status\n")
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
		fmt.Printf("Marshal error: %v\n", err)
		return err
	}
	err = conn.WriteMessage(websocket.BinaryMessage, PakoDeflate(dataStr))
	if err != nil {
		fmt.Printf("Verify error: %v\n", err)
		return err
	}
	return nil
}
