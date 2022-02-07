package service

type Task interface {
	Pull() error
	Push() error
}
