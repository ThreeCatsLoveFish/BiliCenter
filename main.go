package main

import (
	"sub_center/config"
	"sub_center/push"
)

func init() {
	config.LoadConfig()
	push.LoadConfig()
	// TODO: Initialize other packages
}

func main() {
	println("Hello world!")
}
