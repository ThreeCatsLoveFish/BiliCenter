package pull

import (
	"subcenter/application/medalhelper/service"
	"subcenter/domain/push"
	"subcenter/infra/conf"
	"sync"
)

type MedalPull struct{}

func (pull MedalPull) init() []service.Account {
	accounts := make([]service.Account, 0, len(conf.BiliConf.Users))
	for _, user := range conf.BiliConf.Users {
		account := service.NewAccount(user)
		accounts = append(accounts, account)
	}
	return accounts
}

func (pull MedalPull) Obtain() ([]push.Data, error) {
	accounts := pull.init()
	wg := sync.WaitGroup{}
	for _, account := range accounts {
		wg.Add(1)
		go func(user service.Account, wg *sync.WaitGroup) {
			if status := user.Init(); status {
				user.Start(wg)
			}
		}(account, &wg)
	}
	wg.Wait()
	return nil, nil
}
