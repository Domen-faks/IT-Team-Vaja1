package Logic

import (
	"IT-Team-Vaja1/DataStructures"
)

func (c *Controller) GetTasks() (tasks []DataStructures.Task, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetAllTasks()

}

func (c *Controller) GetTaskById(taskId int) (task DataStructures.Task, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetTaskById(taskId)

}

func (c *Controller) PostTask(title string, description string, dateAdded string, workingDateEst string) (task DataStructures.Task, err error) {

	//tukaj bi blo zelo priporočeno preveriti sprejete parametre (stringe), da npr. ne vsebujejo kake sql kode

	return c.db.PostTask(title, description, dateAdded, workingDateEst)

}

func (c *Controller) ChangeTask(id int, title string, description string, dateAdded string, workingDateEst string) (task DataStructures.Task, err error) {

	//tukaj bi blo zelo priporočeno preveriti sprejete parametre (stringe), da npr. ne vsebujejo kake sql kode

	return c.db.ChangeTask(id, title, description, dateAdded, workingDateEst)

}

func (c *Controller) DeleteTask(id int) (task DataStructures.Task, err error) {

	//tukaj bi blo zelo priporočeno preveriti sprejete parametre (stringe), da npr. ne vsebujejo kake sql kode

	return c.db.DeleteTask(id)

}
