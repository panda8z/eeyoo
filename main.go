package main

import (
	"os"

	"github.com/panda8z/eeyoo/model"
	"github.com/panda8z/eeyoo/routes"
)

func main() {
	os.Setenv("GinBlogConfigFile", "./config/config.ini")

	model.InitDb()
	routes.InitRouter()
}
