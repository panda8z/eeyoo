
# Section one  init project

## 1.0 init gitee

- new project
- add name
- add url
- add readme
- add ignore
- add license
- set git-flow pattern

```bash
git clone git@gitee.com:panda8xy/.git
git fetch
git pull
git checkout -b develop origin/develop
```

## 1.1 init go mod

`go mod init github.com/panda8z/eeyoo`

## 1.2 add gin

[Quickstart | Gin Web Framework](https://gin-gonic.com/docs/quickstart/)

`go get -u github.com/gin-gonic/gin`

## 1.3 add main.go and test

```go
package main

import "github.com/gin-gonic/gin"

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080
}
```

## 1.4 project structure

`mkdir api middleware config model routes upload utils web`

## 1.5 add config file tool

[go-ini/ini: 超赞的 Go 语言 INI 文件操作](https://ini.unknwon.io/)

`go get gopkg.in/ini.v1`

## 1.6 add config file

**src/config/config.ini:**

```ini

[app]

[server]
AppMode = debug
HttpPort = :8081

[database]
Db = mysql
DbHost = 127.0.0.1
DbPort = 3307
DbUser = root
DbPassword = 123456
DbName = eeyoo
```

## 1.7 add settings.go

to read config.ini file
**src/utils/settings.go**

```go
package utils

import (
  "log"

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
  file, err := ini.Load("./config/config.ini")
  if err != nil {
    log.Println(err.Error())
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
  DbName = file.Section("server").Key("DbName").MustString("eeyoo")
}

```

## 1.8 add routers.go

to have a common router

**src/routes/router.go:**

```go

package routes

import (
  "net/http"

  "github.com/panda8z/eeyoo/utils"
  "github.com/gin-gonic/gin"
)

func InitRouter() {
  gin.SetMode(utils.AppMode)

  r := gin.Default()

  router := r.Group("api/v1")
  {
    router.GET("hello", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
        "msg": "OK",
      })
    })
  }

  r.Run(utils.HttpPort)

}
```

## 1.9 change the main.go

**src/main.go s**

```go
package main

import (
    "github.com/panda8z/eeyoo/routes"
)

func main() {
    routes.InitRouter()
}

```

## 1.10 Test for first API 

http://localhost:8081/api/v1/hello