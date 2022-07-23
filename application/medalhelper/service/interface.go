package service

import (
	"sync"

	"subcenter/infra/dto"
)

type IConcurrency interface {
	// Exec the action of child and execute retry backup if
	Exec(user Account, work *sync.WaitGroup, child IExec) []dto.MedalInfo
}

type IExec interface {
	// Do represent real action
	Do(user Account, medal dto.MedalInfo) bool
	// Finish represent action complete
	Finish(user Account, medal []dto.MedalInfo)
}

// Action represent a single action for a single user
type IAction interface {
	IConcurrency
	IExec
}
