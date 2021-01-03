package main

import (
	"github.com/panda8z/eeyoo/model"
	"github.com/panda8z/eeyoo/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
