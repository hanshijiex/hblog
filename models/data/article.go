package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"hblog/models"
	"fmt"
	"time"
)

type Article struct{}

var o orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:admin@/blog?charset=utf8")
	o = orm.NewOrm()
	o.Using("default")
}

func(a *Article) Insert(title, content string) (int64) {
	article := new(models.Article)
	article.Content = content
	article.Title = title
	article.Ctime = time.Now().Unix()
	article.Utime = time.Now().Unix()
	id, err := o.Insert(article)
	if err != nil {
		fmt.Println(err)
	}
	return id
}

func(a *Article) List(pageSize, pageNum int32) ([]models.Article)  {
	var articles []models.Article
	num, err := o.Raw("SELECT id, title, content, ctime, utime FROM t_article ORDER BY ctime desc LIMIT ?, ? ",
		pageNum * pageSize, pageNum).QueryRows(&articles)
	if err != nil {
		fmt.Println(err, num)
	}
	fmt.Println(articles)
	return articles
}
