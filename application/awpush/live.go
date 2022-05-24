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

	go func() {
		l.enterRoom(id)
		l.client.Listening()
	}()

	timer := time.NewTimer(5 * time.Minute)
	<-timer.C
	l.client.Close()
	log.Info("Websocket leave room %d", l.roomID)
}

func (l *Live) registerHandlers(handlers ...interface{}) error {
	return websocket.RegisterHandlers(handlers...)
}

// 发送进入房间请求
func (l *Live) enterRoom(roomId int) {
	log.Info("Websocket enter room %d", roomId)
	var err error
	body, _ := json.Marshal(dto.WSEnterRoomBody{
		UID:       12309253,
		RoomID:    roomId,
		ProtoVer:  1,
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
	return func(danmaku *dto.Danmaku) {}
}
