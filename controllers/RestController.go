package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"hblog/models/data"
	"os"
	"hblog/models"
)

type RestController struct {
	beego.Controller
}

var article data.Article

func (r *RestController) Add() {
	title := r.GetString("title")
	content := r.GetString("content")
	var num int64
	num = article.Insert(title, content)
	r.Data["json"] = &num
	r.ServeJSON()
}

func (r *RestController) List() {
	pageSize, err := r.GetInt32("pagesize")
	checkErr(err)
	pageNum, err := r.GetInt32("pagenum")
	checkErr(err)
	var res []models.Article
	res = article.List(pageSize, pageNum)
	r.Data["json"] = &res
	r.ServeJSON()
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}