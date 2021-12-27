package MariaDB

import (
	"IT-Team-Vaja1/DataStructures"
	"errors"
	"fmt"
)

//Funkcija iz interface-a
func (db *MariaDB) GetAllTasks() (tasks []DataStructures.Task, err error) {

	//Naredimo query na bazo
	//Za stavke od katerih ne pričakujemo odgovora (UPDATE, INSERT) uporabimo namesto "Query" "Exec"
	rows, err := db.database.Query("SELECT * from task")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Z defer določimo funkcijo, ki se bo poklicala kadar se funkcija nad njo zaključi.
	//V tem primeru bo to klicano po return

	//Pri SQL-u moremo VEDNO zapreti vrstice.
	//V nasprotnem primeru pride do memory leak-a, vrstice ostanejo odprte, povezave ostaneko odprte
	//in čez čas zmanjka povezav in se dostop do baze ustavi in edina rešitev je potem še reboot
	defer func() {
		err2 := rows.Close()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	}()

	//Ustvarimo objekt Task in neomejen array tipa Task
	var task DataStructures.Task
	tasks = make([]DataStructures.Task, 0)

	//Loop čez vse vrstice
	for rows.Next() {

		//Preberemo podatke vrstice v objekt
		err = rows.Scan(&task.Id, &task.Title, &task.Desc, &task.DateAdded, &task.WorkingDateEst)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//Dodamo objekt task na konec arraya tasks
		tasks = append(tasks, task)

	}

	//Vrnemo podatke. Ne rabimo napisat katero spremenjljivko vrnemo saj je ta definirana in inicializirana na vrhu funkcije
	return

}

func (db *MariaDB) GetTaskById(taskId int) (task DataStructures.Task, err error) {

	rows, err := db.database.Query("SELECT * from task WHERE task_id = ? LIMIT 1", taskId)
	//rows, err := db.database.Query("SELECT * from task")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		err2 := rows.Close()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	}()

	//Prestavimo se na naslednjo vrstico (na prvo vrnjeno vrstico), če je ni pomeni da je odgovor prazen in ne obstaja
	if !rows.Next() {
		err = errors.New("task does not exist")
		return
	}

	err = rows.Scan(&task.Id, &task.Title, &task.Desc, &task.DateAdded, &task.WorkingDateEst)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return

}

func (db *MariaDB) PostTask(title string, description string, dateAdded string, workingDateEst string) (task DataStructures.Task, err error) {

	//fmt.Println(title + " - " + description + " - " + dateAdded + " - " + workingDateEst)
	rows, err := db.database.Query("INSERT INTO task VALUES (NULL, ?,?,?,?)", title, description, dateAdded, workingDateEst)
	//rows, err := db.database.Query("SELECT * from task")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		err2 := rows.Close()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	}()

	return
}

func (db *MariaDB) ChangeTask(id int, title string, description string, dateAdded string, workingDateEst string) (task DataStructures.Task, err error) {

	//rows, err := db.database.Query("SELECT * from task")
	//rows, err := db.database.Query("INSERT INTO task VALUES (NULL, ?,?,?,?)", title, description, dateAdded, workingDateEst)
	rows, err := db.database.Query("UPDATE task SET title = ?, description = ?, date_added = ?, working_date_estimate = ? WHERE task_id = ?", title, description, dateAdded, workingDateEst, id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		err2 := rows.Close()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	}()

	return
}

func (db *MariaDB) DeleteTask(id int) (task DataStructures.Task, err error) {

	//rows, err := db.database.Query("SELECT * from task")
	//rows, err := db.database.Query("INSERT INTO task VALUES (NULL, ?,?,?,?)", title, description, dateAdded, workingDateEst)
	rows, err := db.database.Query("DELETE FROM task WHERE task_id = ?", id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		err2 := rows.Close()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	}()

	return
}
