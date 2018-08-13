package main

import (
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

const (
	host     = "192.168.99.100"
	port     = 32769
	user     = "postgres"
	password = "5763"
	dbName   = "test"
)

func getDBEngine() *xorm.Engine {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	//格式
	engine, err := xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	engine.ShowSQL() //菜鸟必备

	err = engine.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("connect postgresql success")
	return engine
}

//查询所有
func selectAll() {
	var b []Birds
	engine := getDBEngine()
	engine.SQL(`SELECT * FROM "birds" `).Find(&b)
	fmt.Println(b)
}

//条件查询
func selectBird(id int) {
	var b Birds
	engine := getDBEngine()
	engine.Where(`id =? `, id).Find(&b)
	fmt.Println("条件查询", b)
}

//可以用Get查询单个元素
func selectOne(id int) {
	var b Birds
	engine := getDBEngine()
	engine.Id(id).Get(&b)
	//engine.Alias("u").Where("u.id=?",id).Get(&user)
	fmt.Println(b)
}

//

//添加
func InsertUser(b *Birds) bool {
	engine := getDBEngine()
	rows, err := engine.Insert(b)
	if err != nil {
		log.Println(err)
		return false
	}
	if rows == 0 {
		return false
	}
	return true
}

//
func main() {
	// selectAll()
	// selectBird(1)

	var b Birds
	b.Id = 3
	b.Bird = "滴滴滴"
	b.Description = "111111"
	InsertUser(&b)
}

type Birds struct {
	Id          int
	Bird        string
	Description string
}
