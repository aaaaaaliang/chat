package main

import (
	"chat/router"
	"chat/utils"
	"fmt"
)

func main() {
	//初始化路径  数据库
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()
	err := r.Run()
	if err != nil {
		fmt.Println("main.go", err)
		return
	}
}
