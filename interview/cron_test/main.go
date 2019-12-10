package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {

	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		//mysqlDb := db.NewMysqlDb()
		//defer mysqlDb.Close()
		//
		//mysqlDb.Delete(lock.MysqlLock{}, "TIMESTAMPDIFF(SECOND,update_time,NOW()) > 5")
		fmt.Println(1)
	})
	c.Start()

	select {}
}
