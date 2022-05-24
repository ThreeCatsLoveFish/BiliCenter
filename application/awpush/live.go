package awpush

import (
	"encoding/json"
	"strconv"
	"subcenter/infra/log"
	"time"

	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/resource"
	"github.com/botplayerneo/bili-live-api/websocket"
)

type Live struct {
	client *websocket.Client
	roomID int
}

func newLive(roomID int) *Live {
	return &Live{
		roomID: roomID,
	}
}

func (l *Live) start() {
	l.client = websocket.New()
	l.listen()
}

func (l *Live) listen() {
	id, err := resource.RealRoomID(l.roomID)
	if err != nil {
		log.Error("Get room id fail: %v", err)
		return
	}

	if err := l.client.Connect(); err != nil {
		log.Error("Connect ws fail: %v", err)
		return
	}

	go l.enterRoom(id)

	if err := l.client.Listening(); err != nil {
		log.Error("Listen ws fail: %v", err)
		return
	}

	timer := time.NewTimer(5 * time.Minute)
	<-timer.C
	l.client.Close()
}

// RegisterHandlers 注册不同的事件处理
// handler类型需要是定义在 websocket/handler_registration.go 中的类型，如:
// - websocket.DanmakuHandler
// - websocket.GiftHandler
// - websocket.GuardHandler
func (l *Live) registerHandlers(handlers ...interface{}) error {
	return websocket.RegisterHandlers(handlers...)
}

// 发送进入房间请求
func (l *Live) enterRoom(id int) {
	log.Info("Websocket Enter room %d", id)
	// 忽略错误
	var err error
	body, _ := json.Marshal(dto.WSEnterRoomBody{
		RoomID:    id, // 真实房间ID
		ProtoVer:  1,  // 填1
		Platform:  "web",
		ClientVer: "1.6.3",
	})
	if err = l.client.Write(&dto.WSPayload{
		ProtocolVersion: dto.JSON,
		Operation:       dto.RoomEnter,
		Body:            body,
	}); err != nil {
		log.Error("Send enter room ws error: %v", err)
		return
	}
}

func listenRoom(roomId string) {
	room, err := strconv.ParseInt(roomId, 10, 64)
	if err != nil {
		return
	}

	l := newLive(int(room))
	l.registerHandlers(
		danmakuHandler(),
	)
	l.start()
}

func danmakuHandler() websocket.DanmakuHandler {
	return func(danmaku *dto.Danmaku) {
		log.Info("%s: %s\n", danmaku.Uname, danmaku.Content)
	}
}
