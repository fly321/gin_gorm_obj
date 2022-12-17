package main

import (
	"demoProject/global"
	"demoProject/routers"
	"strconv"
)

func main() {
	r := routers.Router()
	r.Run(":" + strconv.Itoa(global.ServerSetting.HttpPort)) // listen and serve on 0.0.0.0:8080
}
