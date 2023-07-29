package models

import (
	"fmt"
	"time"
)
import (
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	Id          int64     `json:"auto"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
	Done        bool      `json:"done"`
}

// 实现自定义表名

func (u *Todo) TableName() string {
	return "todo"
}

func init() {

	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Print("ORM Register Failed")
		return
	}
	err = orm.RegisterDataBase("default", "mysql", "root:ljl123@tcp(127.0.0.1:3306)/GolangDemo?charset=utf8")
	if err != nil {
		fmt.Println("Data base Connection Failed")
		return
	}

	orm.RegisterModel(new(Todo))
	orm.Debug = true
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Informational("[orm] Create table err : ", err)
	}
}
