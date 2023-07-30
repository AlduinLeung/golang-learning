package routers

import (
	"TodoList/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	todoControllers := &controllers.TodoController{}
	beego.Router("/todo", todoControllers, "get:GetAllTodoList")        // get all todos
	beego.Router("/todo", todoControllers, "post:CreateATodo")          // create a todo
	beego.Router("/todo/:id", todoControllers, "get:GetTodoById")       // get todo by id
	beego.Router("/todo/:id", todoControllers, "delete:DeleteTodoById") // delete todo by id
	beego.Router("todo/:id", todoControllers, "put:UpdateTodoById")     // update todo by id
}
