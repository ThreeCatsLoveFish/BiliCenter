package service

import (
	"time"

	"subcenter/application/medalhelper/manager"
	"subcenter/infra/dto"
)

// Like implement IExec, async like 3 times
type ALike struct {
	AsyncAction
}

func (ALike) Do(account Account, medal dto.MedalInfo) bool {
	return manager.LikeInteract(account.user, medal.RoomInfo.RoomID)
}

func (ALike) Finish(account Account, medal []dto.MedalInfo) {
	if len(medal) == 0 {
		account.info("点赞完成")
	} else {
		account.info("点赞未完成,剩余(%d/%d)", len(medal), len(account.medalsLow))
	}
}

// Danmaku implement IExec, default sync, include sending daily danmu
type Danmaku struct {
	SyncAction
}

func (Danmaku) Do(account Account, medal dto.MedalInfo) bool {
	if ok := manager.WearMedal(account.user, medal.Medal.MedalID); !ok {
		return false
	}
	if ok := manager.SendDanmaku(account.user, medal.RoomInfo.RoomID); !ok {
		return false
	}
	time.Sleep(6 * time.Second)
	account.info("%s 房间弹幕打卡完成", medal.AnchorInfo.NickName)
	return true
}

func (Danmaku) Finish(account Account, medal []dto.MedalInfo) {
	if len(medal) == 0 {
		account.info("弹幕打卡完成")
	} else {
		account.info("弹幕打卡未完成,剩余(%d/%d)", len(medal), len(account.medalsLow))
	}
	if account.wearMedal == dto.DefaultMedal {
		manager.TakeoffMedal(account.user)
		account.info("脱下勋章恢复原样")
	} else {
		manager.WearMedal(account.user, account.wearMedal.Medal.MedalID)
		account.info("重新佩戴勋章 %s", account.wearMedal.Medal.MedalName)
	}
}

// WatchLive implement IExec, default async, include sending heartbeat
type WatchLive struct {
	AsyncAction
}

func (WatchLive) Do(account Account, medal dto.MedalInfo) bool {
	times := 80
	for i := 0; i < times; i++ {
		if ok := manager.Heartbeat(
			account.user,
			account.uuid,
			medal.RoomInfo.RoomID,
			medal.Medal.TargetID,
		); !ok {
			return false
		}
		account.info("%s 房间心跳包已发送(%d/%d)", medal.AnchorInfo.NickName, i+1, times)
		time.Sleep(1 * time.Minute)
	}
	return true
}

func (WatchLive) Finish(account Account, medal []dto.MedalInfo) {
	if len(medal) == 0 {
		account.info("每日80分钟完成")
	} else {
		account.info("每日80分钟未完成,剩余(%d/%d)", len(medal), len(account.medalsLow))
	}
}
