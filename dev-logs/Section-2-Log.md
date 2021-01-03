# Section 2 Database and models

## Open MySQL Database and create database

```bash
mysql -uroot -p
create database `eeyoo` character set utf8mb4 collate utf8mb4_unicode_ci;
docker run 
-p 3306:3306 
--name mysql 
-v /root/mysql/conf:/etc/mysql/conf.d 
-v /root/mysql/logs:/logs 
-v /root/mysql/data:/var/lib/mysql 
-e MYSQL_ROOT_PASSWORD=123456 
-d mysql:5.7.32
```

## new models

`cd model && touch user.go article.go category.go db.go`

## add sql driver

[go-sql-driver/mysql: Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package](https://github.com/go-sql-driver/mysql)

`go get -u github.com/go-sql-driver/mysql`

[GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/zh_CN/)

`go get -u gorm.io/gorm`

## init db.go

**src/model/db.go:**

```go
package model

import (
    "fmt"
    "log"
    "time"

    "github.com/panda8z/eeyoo/utils"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

var (
    db  *gorm.DB
    err error
)

func InitDb() {
    db, err := gorm.Open(utils.Db,
        fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
            utils.DbUser,
            utils.DbPassword,
            utils.DbHost,
            utils.DbPort,
            utils.DbName,
        ))

    if err != nil {
        log.Println("database connect err:", err.Error())
    }

    db.SingularTable(true)

    db.AutoMigrate()

    // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
    db.DB().SetMaxIdleConns(10)

    // SetMaxOpenConns 设置打开数据库连接的最大数量。
    db.DB().SetMaxOpenConns(100)

    // SetConnMaxLifetime 设置了连接可复用的最大时间。
    db.DB().SetConnMaxLifetime(time.Hour)
}

func DbClose() {
    db.Close()
}

```

## init user model

**src/model/user.go:**

```go
package model

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"type:varchar(20)" json:"username"`
    Password string `gorm:"type:varchar(20)" json:"password"`
    Role     int    `gorm:"type:int" json:"role"`
}

```

## init category model

**src/model/category.go:**

```go
package model

import "github.com/jinzhu/gorm"

type Category struct {
    gorm.Model
    Name string `gorm:"type:varchar(20);not null;" json:"name"`
}

```

## init article model

**src/model/article.go:**

```go
package model

import "github.com/jinzhu/gorm"

type Article struct {
    gorm.Model
    Title   string `gorm:"type:varchar(100);not null" json:"title"`
    Cid     int    `gorm:"type:int;not null" json:"cid"`
    Desc    string `gorm:"type:varchar(200)" json:"desc"`
    Content string `gorm:"type:longtext" json:"content"`
    Img     string `gorm:"type:varchar(100)" json:"img"`
}

```

## auto migrate

**src/main.go:**

```go
    db.SingularTable(true)

    db.AutoMigrate(&User{}, &Article{}, &Category{})
```

## call in main.go

```go

package main

import (
    "github.com/panda8z/eeyoo/model"
    "github.com/panda8z/eeyoo/routes"
)

func main() {
    model.InitDb()
    routes.InitRouter()
}

```

## open docker start mysql

docker started
mysql container started in 3307

## Run and Test

when you  `go run main.go`
service will start and the mysql database will create new tables.
