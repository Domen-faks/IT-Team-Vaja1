package main

import (
	"IT-Team-Vaja1/API"
	"github.com/gin-gonic/gin"
)

//Naredimo Router objekt da na njega obesimo funkcije
type Router struct {
	engine *gin.Engine
	api    API.Controller
}

func (r *Router) registerRoutes() (err error) {

	//Pot /api/v1
	api := r.engine.Group("/api/v1")

	//Pot /api/v1/todo
	todo := api.Group("/todo")
	r.registerTaskRoutes(todo)

	return

}

func (r *Router) registerTaskRoutes(task *gin.RouterGroup) {

	//Pot /api/v1/task
	task.GET("/", r.api.GetTasks)

	task.POST("/:title/:description/:date_added/:working_date_estimate", r.api.PostTask)

	//Pot /api/v1/task/:todo_id
	task.GET("/:todo_id", r.api.GetTaskById)

	task.PUT("/:todo_id/:title/:description/:date_added/:working_date_estimate", r.api.ChangeTask)

	task.DELETE("/:todo_id", r.api.DeleteTask)
}
