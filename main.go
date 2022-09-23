package main

import (
	"demoProject/routers"
)

func main() {
	r := routers.Router()
	r.Run() // listen and serve on 0.0.0.0:8080
}
