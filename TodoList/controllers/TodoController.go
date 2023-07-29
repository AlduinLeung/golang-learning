package controllers

import (
	"TodoList/models"
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/server/web"
	"strconv"
	"time"
)

type TodoController struct {
	web.Controller
}

// 包外调用，方法名要大写
func (c *TodoController) GetAllTodoList() {
	// create orm object
	c.Ctx.WriteString("All todos are here")

}

func (c *TodoController) CreateATodo() {
	o := orm.NewOrm()
	todo := models.Todo{Id: 1, Title: "First Todo", Description: "Test Database", CreatedAt: time.Now(), UpdateAt: time.Now(), Done: false}
	code, err := o.Insert(&todo)
	if err != nil {
		logs.Informational("Insert Failed", code)
	}
	c.Ctx.WriteString("Your Todos")
}

func (c *TodoController) GetTodoById() {
	// 获取请求参数
	params := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(params, 10, 64)
	o := orm.NewOrm()
	todo := models.Todo{Id: id}
	err = o.Read(&todo)
	if err != nil {
		logs.Informational("Read Failed", err)
	}
	if todo.Id != 0 && err == nil {
		logs.Informational(todo)
		c.Data["json"] = todo
		c.ServeJSON()
	} else {
		c.Ctx.WriteString("Todo Not Found")
	}
}

func (c *TodoController) DeleteTodoById() {
	params := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(params, 10, 64)
	o := orm.NewOrm()
	todo := models.Todo{Id: id}
	_, err = o.Delete(&todo)
	if err != nil {
		logs.Informational("Delete Failed", err)
	}
	c.Ctx.WriteString("Delete Success")
}
func (c *TodoController) UpdateTodoById() {

}

func main() {
	web.AutoRouter(&TodoController{})
	web.Run()
}
