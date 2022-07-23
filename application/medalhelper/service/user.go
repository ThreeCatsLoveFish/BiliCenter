package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"subcenter/application/medalhelper/manager"
	"subcenter/domain/push"
	"subcenter/infra/conf"
	"subcenter/infra/dto"
	"subcenter/infra/log"

	"github.com/TwiN/go-color"
	"github.com/google/uuid"
	"github.com/sethvargo/go-retry"
	"github.com/tidwall/gjson"
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
	format = color.Green + "[INFO] " + color.Reset + format
	format = color.Reset + color.Blue + account.Name + color.Reset + " " + format
	log.PrintColor(format, v...)
}

func (account *Account) loginVerify() bool {
	resp, err := manager.LoginVerify(account.user)
	if err != nil || resp.Data.Mid == 0 {
		account.isLogin = false
		return false
	}
	account.Uid = resp.Data.Mid
	account.Name = resp.Data.Name
	account.isLogin = true
	account.info("登录成功")
	return true
}

func (account *Account) signIn() error {
	signInfo, err := manager.SignIn(account.user)
	if err != nil {
		return nil
	}
	resp := gjson.Parse(signInfo)
	if resp.Get("code").Int() == 0 {
		signed := resp.Get("data.hadSignDays").String()
		all := resp.Get("data.allDays").String()
		account.info("签到成功, 本月签到次数: %s/%s", signed, all)
	} else {
		account.info("%s", resp.Get("message").String())
	}

	userInfo, err := manager.GetUserInfo(account.user)
	if err != nil {
		return nil
	}
	level := userInfo.Data.Exp.UserLevel
	unext := userInfo.Data.Exp.Unext
	account.info("当前用户UL等级: %d, 还差 %d 经验升级", level, unext)
	return nil
}

func (account *Account) setMedals() {
	// Clean medals storage
	account.medalsLow = make([]dto.MedalInfo, 0, 10)
	account.remainMedals = make([]dto.MedalInfo, 0, 10)
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
		account.signIn()
		account.setMedals()
		return true
	} else {
		pushEnd := push.NewPush(account.user.Push)
		msg := fmt.Sprintf("用户登录失败, cookie: %s", account.user.Cookie)
		pushEnd.Submit(push.Data{
			Title:   "# 用户登录失败",
			Content: msg,
		})
		return false
	}
}

func (account *Account) RunOnce() bool {
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
		backOff := retry.NewConstant(5 * time.Second)
		backOff = retry.WithMaxRetries(3, backOff)
		retry.Do(context.Background(), backOff, func(ctx context.Context) error {
			if ok := account.RunOnce(); !ok {
				return retry.RetryableError(errors.New("task not complete"))
			}
			return nil
		})
		account.report()
	} else {
		log.Error("用户未登录, cookie: %s", account.user.Cookie)
	}
	wg.Done()
}
