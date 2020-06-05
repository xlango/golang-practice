package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/gommon/log"
	"time"
)

type User struct {
	//gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num      int    `gorm:"AUTO_INCREMENT"` // 自增
}

func main() {
	//db, err := gorm.Open("sqlite3", "sqlite/gorm.db?_auth_user=admin&_auth_pass=admin&_auth_crypt=sha1")
	db, err := gorm.Open("sqlite3", "file:sqlite/test.s3db?_auth&_auth_user=admin&_auth_pass=admin&_auth_crypt=sha1")
	if err != nil {
		log.Error(err.Error())
	}
	defer db.Close()

	// 为模型`User`创建表
	db.CreateTable(&User{})

	// 创建表`users'时将“ENGINE = InnoDB”附加到SQL语句
	//db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	db.NewRecord(user) // => 主键为空返回`true`

	db.Create(&user)

	db.NewRecord(user) // => 创建`user`后返回`false`
}
