package awpush

import (
	"fmt"
	"subcenter/manager"

	"github.com/gorilla/websocket"
)

// ping send ping message
func ping(conn *websocket.Conn) error {
	err := conn.WriteMessage(
		websocket.TextMessage,
		[]byte("ping"),
	)
	if err != nil {
		fmt.Printf("Ping error: %v\n", err)
	}
	return err
}

// pong receive pong message
func pong(conn *websocket.Conn) error {
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

// verify used for connect awpush
func verify(conn *websocket.Conn, uid, token string) error {
	err := conn.WriteMessage(
		websocket.BinaryMessage,
		NewVerify(uid, token),
	)
	if err != nil {
		fmt.Printf("Verify error: %v\n", err)
		return err
	}
	return nil
}

func establish() (ws *websocket.Conn, err error) {
	conn, _, err := websocket.DefaultDialer.Dial(biliConfig.Wss, nil)
	if err != nil {
		fmt.Printf("Dial error: %v\n", err)
		return nil, err
	}
	_, msg, err := conn.ReadMessage()
	if err != nil {
		fmt.Printf("ReadMessage error: %v\n", err)
		return nil, err
	}
	res := string(manager.PakoInflate(msg))
	greet := `{"code":0,"type":"WS_OPEN","data":"连接成功"}`
	if res != greet {
		fmt.Printf("Greeting error, obtain: %s\n", res)
		return nil, fmt.Errorf("greeting error, obtain: %s", res)
	}
	err = ping(conn)
	if err != nil {
		fmt.Printf("Ping failed, error: %v", err)
	}
	err = pong(conn)
	if err != nil {
		fmt.Printf("Pong failed, error: %v", err)
	}
	err = verify(conn, biliConfig.Uid, biliConfig.Token)
	if err != nil {
		fmt.Printf("Verify failed, error: %v", err)
	}
	fmt.Println("Verify success")
	return conn, err
}
