package main

import (
	_ "pybbs-go/routers"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "pybbs-go/models"
    _ "github.com/go-sql-driver/mysql"
    _ "pybbs-go/utils"
    _ "pybbs-go/templates"
)

func init(){
    orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("jdbc.username") + ":" + beego.AppConfig.String("jdbc.password") + "@/pybbs-go?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai", 30)
    orm.RegisterModel(
		new(models.User),
		new(models.Topic),
		new(models.Section),
		new(models.Reply),
		new(models.ReplyUpLog),
        new(models.Role),
        new(models.Permission))
    orm.RunSyncdb("default", false, true)
}

func main() {
    //orm.Debug = true
    //ok, err := regexp.MatchString("/topic/edit/[0-9]+", "/topic/edit/123")
    //beego.Debug(ok, err)
	beego.Run()
}

