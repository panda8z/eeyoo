package utils

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	// file, err := ini.Load("./config/config.ini")
	os.Setenv("CONFIG_FILE", "/Users/zcj/panda/git4me/gin-blog/config/config.ini")

	file, err := ini.Load(os.Getenv("CONFIG_FILE"))
	if err != nil {
		log.Println(err.Error() + os.Getenv("CONFIG_FILE"))
		log.Println("Config file not found or open it with error, pleas check src/config/config.ini ")
		panic(err)
	}
	loadServer(file)
	loadDatabase(file)

}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("8081")
}

func loadDatabase(file *ini.File) {
	Db = file.Section("server").Key("Db").MustString("mysql")
	DbHost = file.Section("server").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("server").Key("DbPort").MustString("3307")
	DbUser = file.Section("server").Key("DbUser").MustString("root")
	DbPassword = file.Section("server").Key("DbPassword").MustString("123456")
	DbName = file.Section("server").Key("DbName").MustString("ginblog")
}
