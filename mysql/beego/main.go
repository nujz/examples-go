package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Children struct {
	Id       int
	Username string `orm:"size(64)"`
}

func beegoInitMySQL() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "test@/test?charset=utf8", 30)
	orm.RegisterModel(new(Children))

	orm.RunSyncdb("default", false, true)
}

func main() {
	beegoInitMySQL()

	o := orm.NewOrm()
	user := Children{Username: "tom"}

	id, err := o.Insert(&user)
	fmt.Println(id, err)

	u := Children{Id: user.Id}
	if err := o.Read(&u); err != nil {
		fmt.Println(err)
	}

	user.Username = "bob"
	num, err := o.Update(&user)
	fmt.Println(num, err)

	num, err = o.Delete(&user)
	fmt.Println(num, err)
}
