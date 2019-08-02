package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:12345678@/xrt_open?charset=utf8")
	fmt.Println(err)
	fmt.Println(engine.Ping())
}
