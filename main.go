package main

import (
	"os"

	"gitee.com/panda8xy/gin-blog/model"
	"gitee.com/panda8xy/gin-blog/routes"
)

func main() {
	os.Setenv("GinBlogConfigFile", "/Users/zcj/panda/git4me/gin-blog/config/config.ini")

	model.InitDb()
	routes.InitRouter()
}
