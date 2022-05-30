package awpush

import (
	"fmt"
	"subcenter/domain/push"
	"subcenter/infra"
	"subcenter/infra/log"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

type AWPushClient struct {
	// Basic connection
	conn *websocket.Conn // websocket connection

	// Counter for reporting number of lottery
	recv int32 // received lottery number
	join int32 // joined lottery number

	// Timer used for trigger action
	report  *log.DayTicker // report lottery status
	reset   *time.Timer    // reconnect awpush server
	sleep   *time.Timer    // sleep before handle next message
	timeout *time.Timer    // send heartbeat
}

func NewAWPushClient() AWPushClient {
	return AWPushClient{
		conn:    nil,
		report:  log.NewDayTicker(),
		timeout: time.NewTimer(time.Second * 30),
		reset:   time.NewTimer(time.Microsecond),
		sleep:   time.NewTimer(time.Second),
	}
}

// establish create a new websocket connection
func establish() (ws *websocket.Conn, err error) {
	conn, _, err := websocket.DefaultDialer.Dial(biliConfig.Wss, nil)
	if err != nil {
		log.Error("Dial error: %v", err)
		return nil, err
	}
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Error("ReadMessage error: %v", err)
		return nil, err
	}
	res := string(infra.PakoInflate(msg))
	greet := `{"code":0,"type":"WS_OPEN","data":"连接成功"}`
	if res != greet {
		log.Error("Greeting error, obtain: %s", res)
		return nil, fmt.Errorf("greeting error, obtain: %s", res)
	}
	err = infra.Ping(conn)
	if err != nil {
		log.Error("Ping failed, error: %v", err)
	}
	err = infra.Verify(conn, biliConfig.Uid, biliConfig.Token)
	if err != nil {
		log.Error("Verify failed, error: %v", err)
	}
	log.Info("AwPush Verify sent")
	return conn, err
}

func (tc *AWPushClient) Run() {
	var err error
	for {
		select {
		case <-tc.reset.C:
			if tc.conn != nil {
				tc.conn.Close()
				tc.conn = nil
			}
			if tc.conn, err = establish(); err != nil {
				log.Error("Establish failed, error: %v", err)
				push.NewPush("threecats").Submit(push.Data{
					Title:   "# AWpush establish failed",
					Content: err.Error(),
				})
				continue
			}
			log.Debug("Reconnect success")
		case <-tc.timeout.C:
			if err = infra.Ping(tc.conn); err != nil {
				log.Error("send heartbeat error: %v", err)
				tc.reset.Reset(time.Microsecond)
				continue
			}
			log.Debug("Heartbeat sent")
		case <-tc.sleep.C:
			if err = HandleMsg(tc); err != nil {
				log.Error("handle failed, error: %v", err)
				tc.reset.Reset(time.Microsecond)
				push.NewPush("threecats").Submit(push.Data{
					Title:   "# Handle awpush msg error",
					Content: err.Error(),
				})
			}
		case <-tc.report.C:
			push.NewPush("threecats").Submit(push.Data{
				Title:   "# awpush report",
				Content: fmt.Sprintf("Recv: %d, join: %d", tc.recv, tc.join),
			})
			atomic.StoreInt32(&tc.recv, 0)
			atomic.StoreInt32(&tc.join, 0)
		}
	}
}
