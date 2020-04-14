// mgotest project main.go
package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `bson:"_id"`
	Username  string        `bson:"name"`
	Pass      string        `bson:"pass"`
	Regtime   int64         `bson:"regtime"`
	Interests []string      `bson:"interests"`
}

const URL string = "39.108.147.36:27017"

var c *mgo.Collection
var session *mgo.Session

func (user User) ToString() string {
	return fmt.Sprintf("%#v", user)
}

func init() {
	session, _ = mgo.Dial(URL)
	//切换到数据库
	db := session.DB("test")
	//切换到collection
	c = db.C("mgotest")
}

//新增数据
func add() {
	//    defer session.Close()
	stu1 := new(User)
	stu1.Id = bson.NewObjectId()
	stu1.Username = "stu1_name"
	stu1.Pass = "stu1_pass"
	stu1.Regtime = time.Now().Unix()
	stu1.Interests = []string{"象棋", "游泳", "跑步"}
	err := c.Insert(stu1)
	if err == nil {
		fmt.Println("插入成功")
	} else {
		fmt.Println(err.Error())
		defer panic(err)
	}
}

//查询
func find() {
	//    defer session.Close()
	var users []User
	//    c.Find(nil).All(&users)
	c.Find(bson.M{"name": "stu1_name"}).All(&users)
	for _, value := range users {
		fmt.Println(value.ToString())
	}
	//根据ObjectId进行查询
	idStr := "5e956c75691a43116c7b7b21"
	objectId := bson.ObjectIdHex(idStr)
	user := new(User)
	c.Find(bson.M{"_id": objectId}).One(user)
	fmt.Println(user)
}

//根据id进行修改
func update() {
	interests := []string{"象棋", "游泳", "跑步"}
	err := c.Update(bson.M{"_id": bson.ObjectIdHex("5e956c75691a43116c7b7b21")}, bson.M{"$set": bson.M{
		"name":      "修改后的name",
		"pass":      "修改后的pass",
		"regtime":   time.Now().Unix(),
		"interests": interests,
	}})
	if err != nil {
		fmt.Println("修改失败")
	} else {
		fmt.Println("修改成功")
	}
}

//删除
func del() {
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex("5e956c75691a43116c7b7b21")})
	if err != nil {
		fmt.Println("删除失败" + err.Error())
	} else {
		fmt.Println("删除成功")
	}
}
func main() {
	//add()
	//find()
	//update()
	del()
}
