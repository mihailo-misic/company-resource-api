package main

import (
	_ "github.com/mihailo-misic/company-resource-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	// Setting up Database and ORM.
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("crm", "postgres", "postgres://postgres:secret@localhost/crm?sslmode=disable&charset=utf8")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
	orm.RunSyncdb("crm", false, false)
}
