package main

import (
	"MyUser_System/config"
	"MyUser_System/router"
)

func Init() {
	config.InitConfig()
}

func main() {
	Init()
	router.InitRouterAndServe()

}
