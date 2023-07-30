package controllers

import (
	"TodoList/Utils"
	"TodoList/models"
	"encoding/json"
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
	o := orm.NewOrm()
	todo := models.Todo{}
	slice := []models.Todo{}
	_, err := o.QueryTable(todo).All(&slice)
	if err != nil {
		res := Utils.Response(500, "Failed", err)
		c.Data["json"] = res
		c.ServeJSON()
	} else {
		res := Utils.Response(200, "Success", slice)
		c.Data["json"] = res
		c.ServeJSON()
	}
}

// Create A todo
func (c *TodoController) CreateATodo() {
	todo := models.Todo{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &todo)
	todo.CreatedAt = time.Now()
	todo.UpdateAt = time.Now()
	logs.Informational(todo)
	o := orm.NewOrm()
	_, err := o.Insert(&todo)
	if err != nil {
		res := Utils.Response(500, "Failed", err)
		c.Data["json"] = res
		c.ServeJSON()
	}
	res := Utils.Response(200, "Success", todo)
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *TodoController) GetTodoById() {
	// 获取请求参数
	params := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(params, 10, 64)
	o := orm.NewOrm()
	todo := models.Todo{Id: id}
	err = o.Read(&todo)
	if err != nil {
		res := Utils.Response(500, "Read Failed", err)
		c.Data["json"] = res
		c.ServeJSON()
	} else {
		res := Utils.Response(200, "Success", todo)
		c.Data["json"] = res
		c.ServeJSON()
	}
}

func (c *TodoController) DeleteTodoById() {
	params := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(params, 10, 64)
	o := orm.NewOrm()
	todo := models.Todo{Id: id}
	_, err = o.Delete(&todo)
	res := Utils.ResponseData{}
	if err != nil {
		res = Utils.Response(500, "Failed", err)
		c.Data["json"] = res
		c.ServeJSON()
		logs.Informational("Delete Failed", err)
	} else {
		res = Utils.Response(200, "Success", todo)
		c.Data["json"] = res
		c.ServeJSON()
	}

}

func (c *TodoController) UpdateTodoById() {
	params := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(params, 10, 64)
	o := orm.NewOrm()
	todo := models.Todo{Id: id}
	err = o.Read(&todo)
	if err != nil {
		res := Utils.Response(500, "Read Failed", err)
		c.Data["json"] = res
		c.ServeJSON()
	} else {
		json.Unmarshal(c.Ctx.Input.RequestBody, &todo)
		todo.UpdateAt = time.Now()
		_, err = o.Update(&todo)
		if err != nil {
			res := Utils.Response(500, "Update Failed", err)
			c.Data["json"] = res
			c.ServeJSON()
		} else {
			res := Utils.Response(200, "Success", todo)
			c.Data["json"] = res
			c.ServeJSON()
		}
	}
}

func main() {
	web.AutoRouter(&TodoController{})
	web.Run()
}
