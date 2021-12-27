package API

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (a *Controller) GetTasks(c *gin.Context) {

	//Dobimo uporabnike - Kličemo Logic in ne direkt baze!
	tasks, err := a.c.GetTasks()
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	//Avtomatsko serializiramo objekt task v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, tasks)

}

func (a *Controller) GetTaskById(c *gin.Context) {

	//Dobimo todo_id tipa string iz naslova in ga pretvorimo v int
	taskId, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	//Dobimo uporabnika glede na ID - Kličemo Logic in ne direkt baze!
	task, err := a.c.GetTaskById(taskId)
	if err != nil {
		//Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	//Avtomatsko serializiramo objekt task v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, task)

}

func (a *Controller) PostTask(c *gin.Context) {
	//Dobimo title tipa string iz naslova
	title := c.Param("title")

	//Dobimo description tipa string iz naslova
	description := c.Param("description")

	//Dobimo date:added tipa string iz naslova
	dateAdded := c.Param("date_added")

	//Dobimo working_date_estimate tipa string iz naslova
	workingDateEst := c.Param("working_date_estimate")

	//Kličemo Logic in ne direkt baze!
	_, err := a.c.PostTask(title, description, dateAdded, workingDateEst)
	if err != nil {
		return
	}

}

func (a *Controller) ChangeTask(c *gin.Context) {

	//Dobimo task_id tipa string iz naslova in ga pretvorimo v int
	id, err := strconv.Atoi(c.Param("todo_id"))

	//Dobimo title tipa string iz naslova
	title := c.Param("title")

	//Dobimo description tipa string iz naslova
	description := c.Param("description")

	//Dobimo date_added tipa string iz naslova
	dateAdded := c.Param("date_added")

	//Dobimo working_date_estimate tipa string iz naslova
	workingDateEst := c.Param("working_date_estimate")

	//Dobimo uporabnika glede na ID - Kličemo Logic in ne direkt baze!
	_, err = a.c.ChangeTask(id, title, description, dateAdded, workingDateEst)
	if err != nil {
		return
	}

}

func (a *Controller) DeleteTask(c *gin.Context) {

	//Dobimo todo_id tipa string iz naslova in ga pretvorimo v int
	taskId, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	//Dobimo uporabnika glede na ID - Kličemo Logic in ne direkt baze!
	_, err = a.c.DeleteTask(taskId)
	if err != nil {
		//Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

}
