package Queries

import "strconv"

const AllTasks string = "SELECT * FROM \"Todolist_Schema\".task"

const OneTask string = "SELECT * FROM \"Todolist_Schema\".task Where \"IdTask\"=$1"

const DeleteTask = `DELETE FROM "Todolist_Schema".task WHERE "IdTask"=$1`

func UpdateTask(taskId int) string {

	id := strconv.Itoa(taskId)

	return `UPDATE "Todolist_Schema".task SET title=$1, description=$2, "createdAt" = $3, deadline= $4,active = $5 WHERE "IdTask"=` + id
}

func InsertTask() string {
	return `INSERT INTO "Todolist_Schema".task (title, description, deadline, "createdAt", active) VALUES ($1, $2, $3, $4, $5)`
}
