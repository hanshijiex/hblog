package models

import "github.com/astaxie/beego/orm"

type Article struct {
	Id int
	Title string
	Content string
	Ctime int64
	Utime int64
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModelWithPrefix("t_", new(Article))
}
