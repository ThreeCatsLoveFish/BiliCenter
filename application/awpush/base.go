package awpush

import (
	"fmt"
	"log"
	"subcenter/domain/push"
	"subcenter/infra"
	"time"

	"github.com/gorilla/websocket"
)

// Establish create a new websocket connection
func Establish() (ws *websocket.Conn, err error) {
	conn, _, err := websocket.DefaultDialer.Dial(biliConfig.Wss, nil)
	if err != nil {
		log.Default().Printf("Dial error: %v", err)
		return nil, err
	}
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Default().Printf("ReadMessage error: %v", err)
		return nil, err
	}
	res := string(infra.PakoInflate(msg))
	greet := `{"code":0,"type":"WS_OPEN","data":"连接成功"}`
	if res != greet {
		log.Default().Printf("Greeting error, obtain: %s", res)
		return nil, fmt.Errorf("greeting error, obtain: %s", res)
	}
	err = infra.Ping(conn)
	if err != nil {
		log.Default().Printf("Ping failed, error: %v", err)
	}
	err = infra.Pong(conn)
	if err != nil {
		log.Default().Printf("Pong failed, error: %v", err)
	}
	err = infra.Verify(conn, biliConfig.Uid, biliConfig.Token)
	if err != nil {
		log.Default().Printf("Verify failed, error: %v", err)
	}
	log.Default().Println("[INFO] AwPush Verify success")
	return conn, err
}

type AWPushClient struct {
	// Basic connection
	conn *websocket.Conn // websocket connection

	// Timer used for trigger action
	timeout *time.Ticker // heartbeat period
	reset   *time.Ticker // reconnect period
	sleep   *time.Timer  // sleep time before execute next task
}

func NewAWPushClient() AWPushClient {
	conn, err := Establish()
	if err != nil {
		log.Default().Printf("establish failed, error: %v", err)
	}
	return AWPushClient{
		conn:    conn,
		timeout: time.NewTicker(time.Second * 30),
		reset:   time.NewTicker(time.Hour * 4),
		sleep:   time.NewTimer(time.Second * 1),
	}
}

func (tc *AWPushClient) Serve() {
	var err error
	for {
		select {
		case <-tc.timeout.C:
			if err = infra.Ping(tc.conn); err != nil {
				log.Default().Printf("send heartbeat error: %v", err)
				continue
			}
			log.Default().Printf("[DEBUG] Heartbeat sent")
		case <-tc.sleep.C:
			if err = HandleMsg(tc.conn, tc.sleep); err != nil {
				log.Default().Printf("handle failed, error: %v", err)
				push.NewPush("threecats").Submit(push.Data{
					Title:   "# Handle awpush msg error",
					Content: err.Error(),
				})
			}
		case <-tc.reset.C:
			if tc.conn, err = Establish(); err != nil {
				log.Default().Printf("Reconnect failed, error: %v", err)
				continue
			}
			log.Default().Printf("[DEBUG] Reconnect success")
		}
	}
}
