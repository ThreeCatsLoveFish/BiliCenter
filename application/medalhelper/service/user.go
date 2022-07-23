package service

import (
	"fmt"
	"sync"
	"time"

	"subcenter/application/medalhelper/manager"
	"subcenter/domain/push"
	"subcenter/infra/conf"
	"subcenter/infra/dto"
	"subcenter/infra/log"

	"github.com/google/uuid"
)

type Account struct {
	// 用户ID
	Uid int
	// 用户名称
	Name string
	// 是否登录
	isLogin bool
	// UUID
	uuid []string

	// 用户配置信息
	user conf.User

	// 用户佩戴的勋章
	wearMedal dto.MedalInfo
	// 用户等级小于20的勋章
	medalsLow []dto.MedalInfo
	// 今日亲密度没满的勋章
	remainMedals []dto.MedalInfo

	// 日志信息
	message string
}

func NewAccount(user conf.User) Account {
	return Account{
		user:      user,
		wearMedal: dto.DefaultMedal,
		uuid:      []string{uuid.NewString(), uuid.NewString()},
		message:   "",
	}
}

func (account Account) info(format string, v ...interface{}) {
	format = account.Name + " " + format
	log.Info(format, v...)
}

func (account *Account) loginVerify() bool {
	userInfo, err := manager.GetUserInfo(account.user)
	if err != nil {
		return false
	}
	account.isLogin = true
	account.Uid = userInfo.Data.UID
	account.Name = userInfo.Data.Uname
	account.info("登录成功")

	if _, err := manager.SignIn(account.user); err == nil {
		account.info("签到成功")
	}
	return true
}

func (account *Account) setMedals() {
	// Clean medals storage
	account.medalsLow = make([]dto.MedalInfo, 0, 10)
	account.remainMedals = make([]dto.MedalInfo, 0, 10)
	// Clean bad cache
	manager.GetMedal(account.user)
	time.Sleep(5 * time.Second)
	// Fetch and update medals
	medals, wearMedal := manager.GetMedal(account.user)
	if wearMedal {
		account.wearMedal = medals[0]
	}
	// Default blacklist
	for _, medal := range medals {
		if medal.RoomInfo.RoomID == 0 {
			continue
		}
		if medal.Medal.Level <= 20 {
			account.medalsLow = append(account.medalsLow, medal)
			if medal.Medal.TodayFeed < 1500 {
				account.remainMedals = append(account.remainMedals, medal)
			}
		}
	}
}

func (account *Account) checkMedals() bool {
	account.setMedals()
	fullMedalList := make([]string, 0, len(account.medalsLow))
	failMedalList := make([]string, 0)
	for _, medal := range account.medalsLow {
		if medal.Medal.TodayFeed == 1500 {
			fullMedalList = append(fullMedalList, medal.AnchorInfo.NickName)
		} else {
			failMedalList = append(failMedalList, medal.AnchorInfo.NickName)
		}
	}
	account.message = fmt.Sprintf(
		"20级以下牌子共 %d 个\n【1500】%d个\n【1500以下】 %v等 %d个\n",
		len(account.medalsLow), len(fullMedalList),
		failMedalList, len(failMedalList),
	)
	account.info(account.message)
	return len(fullMedalList) == len(account.medalsLow)
}

func (account *Account) report() {
	pushEnd := push.NewPush(account.user.Push)
	pushEnd.Submit(push.Data{
		Title:   "# 今日亲密度获取情况如下",
		Content: fmt.Sprintf("用户%s，%s", account.Name, account.message),
	})
}

func (account *Account) Init() bool {
	if account.loginVerify() {
		account.setMedals()
	} else {
		msg := fmt.Sprintf("用户Cookie过期: %s", account.user.Cookie)
		pushEnd := push.NewPush(account.user.Push)
		pushEnd.Submit(push.Data{
			Title:   "# 用户Cookie过期",
			Content: "请主动联系管理员更新Cookie",
		})
		account.info(msg)
		return false
	}
	return true
}

func (account *Account) Run() bool {
	task := NewTask(*account, []IAction{
		&ALike{},
		&Danmaku{},
		&WatchLive{},
	})
	task.Start()
	return account.checkMedals()
}

func (account *Account) Start(wg *sync.WaitGroup) {
	if account.isLogin {
		account.Run()
		account.report()
	} else {
		log.Error("用户未登录, cookie: %s", account.user.Cookie)
	}
	wg.Done()
}
