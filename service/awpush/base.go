package awpush

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type AWPushClient struct {
	// Basic connection
	conn *websocket.Conn // websocket connection

	// Timer used for trigger action
	timeout *time.Ticker // heartbeat period
	sleep   *time.Timer  // sleep time before execute next task
}

func NewAWPushClient() AWPushClient {
	conn, err := establish()
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
			if err := ping(tc.conn); err != nil {
				fmt.Printf("send heartbeat error: %v", err)
			}
		case <-tc.sleep.C:
			if err := handleMsg(tc.conn, tc.sleep); err != nil {
				fmt.Printf("handle failed, error: %v", err)
			}
		}
	}
}
