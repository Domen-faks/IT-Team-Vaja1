package API

import "IT-Team-Vaja1/Logic"

type Controller struct {
	c *Logic.Controller
}

func NewController(c *Logic.Controller) Controller {
	return Controller{c}
}

func (a *Controller) GetLogicController() *Logic.Controller {
	return a.c
}
