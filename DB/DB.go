package DB

import "IT-Team-Vaja1/DataStructures"

type DB interface {
	Init() (err error)

	GetAllTasks() (tasks []DataStructures.Task, err error)
	GetTaskById(id int) (task DataStructures.Task, err error)

	PostTask(title string, description string, dateAdded string, workingDateEst string) (task DataStructures.Task, err error)

	ChangeTask(id int, title string, description string, dateAdded string, workingDateEst string) (task DataStructures.Task, err error)

	DeleteTask(id int) (task DataStructures.Task, err error)
}
