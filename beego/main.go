package main

import (
	"fmt"
_ "nxxlzx/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	if err != nil {
		fmt.Printf("数据库连接失败：%s\n", err.Error())
		return
	}
	// 开启日志记录
	logs.SetLogger("file")
	logs.SetLogger(logs.AdapterFile, `{"filename":"log/run.log","maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
	logs.Async(1e3)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/docs"] = "swagger"
	}
	beego.Run()
}
