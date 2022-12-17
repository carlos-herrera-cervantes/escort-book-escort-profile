package config

import "os"

type app struct {
	Port string
}

var singleApp *app

func InitApp() *app {
	if singleApp != nil {
		return singleApp
	}

	lock.Lock()
	defer lock.Unlock()

	singleApp = &app{
		Port: os.Getenv("PORT"),
	}

	return singleApp
}
