package main

import (
	"gitee.com/panda8xy/gin-blog/model"
	"gitee.com/panda8xy/gin-blog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
