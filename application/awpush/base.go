package awpush

import (
	"fmt"
	"subcenter/infra"
	"time"

	"github.com/gorilla/websocket"
)

// Establish create a new websocket connection
func Establish() (ws *websocket.Conn, err error) {
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
	res := string(infra.PakoInflate(msg))
	greet := `{"code":0,"type":"WS_OPEN","data":"连接成功"}`
	if res != greet {
		fmt.Printf("Greeting error, obtain: %s\n", res)
		return nil, fmt.Errorf("greeting error, obtain: %s", res)
	}
	err = infra.Ping(conn)
	if err != nil {
		fmt.Printf("Ping failed, error: %v", err)
	}
	err = infra.Pong(conn)
	if err != nil {
		fmt.Printf("Pong failed, error: %v", err)
	}
	err = infra.Verify(conn, biliConfig.Uid, biliConfig.Token)
	if err != nil {
		fmt.Printf("Verify failed, error: %v", err)
	}
	fmt.Println("Verify success")
	return conn, err
}

type AWPushClient struct {
	// Basic connection
	conn *websocket.Conn // websocket connection

	// Timer used for trigger action
	timeout *time.Ticker // heartbeat period
	sleep   *time.Timer  // sleep time before execute next task
}

func NewAWPushClient() AWPushClient {
	conn, err := Establish()
	if err != nil {
		fmt.Printf("establish failed, error: %v", err)
	}
	return AWPushClient{
		conn:    conn,
		timeout: time.NewTicker(time.Second * 30),
		sleep:   time.NewTimer(time.Second * 1),
	}
}

func (tc *AWPushClient) Serve() {
	for {
		select {
		case <-tc.timeout.C:
			if err := infra.Ping(tc.conn); err != nil {
				fmt.Printf("send heartbeat error: %v", err)
			}
		case <-tc.sleep.C:
			if err := HandleMsg(tc.conn, tc.sleep); err != nil {
				fmt.Printf("handle failed, error: %v", err)
			}
		}
	}
}
