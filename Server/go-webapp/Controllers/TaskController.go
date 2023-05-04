package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go-webapp/Database"
	"go-webapp/Models"
	"go-webapp/Queries"
	"log"
	"net/http"
	"strconv"
)

var db, err = Database.Connection()

func GetTasks(w http.ResponseWriter, r *http.Request) {

	if err != nil {
		log.Fatalf("Can't connect to database\nError:%s", err)
	}
	var tasks []Models.Task
	var task Models.Task
	rows, err := db.Query(Queries.AllTasks)
	if err != nil {
		log.Fatalf("Incorrect Query\nError:%s", err)
	}

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.Deadline, &task.Active)

		if err != nil {
			log.Fatalf("An error ocurred!\nError: %v", err)
		}
		tasks = append(tasks, task)
	}
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(tasks)

	fmt.Printf("A %v was made: %v", r.Method, r.URL)

}

func GetTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	id := chi.URLParam(r, "id")

	TaskID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	row := db.QueryRow(Queries.OneTask, TaskID)

	var task Models.Task
	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.Deadline, &task.Active)
	if err != nil {
		http.Error(w, "Failed to get task", http.StatusInternalServerError)
		return
	}

	fmt.Printf("A %v was made: %v", r.Method, r.URL)

	json.NewEncoder(w).Encode(task)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	var task Models.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		log.Fatalf("An error has ocurred!\nError: %v", err)
	}

	db.Exec(Queries.InsertTask(), task.Title, task.Description, task.CreatedAt, task.Deadline, task.Active)

	fmt.Printf("A %v was made: %v", r.Method, r.URL)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	if err != nil {
		log.Fatalf("Cannot connect to database %v", err)
	}

	id := chi.URLParam(r, "id")

	_, err := db.Exec(Queries.DeleteTask, id)
	if err != nil {
		log.Fatalf("Invalid id %v", err)
	}
	fmt.Printf("A %v was made: %v", r.Method, r.URL)

}

func PutTask(w http.ResponseWriter, r *http.Request) {

	var task Models.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		log.Fatalf("Cannot connect to database %v", err)
	}

	id := chi.URLParam(r, "id")

	idTask, err := strconv.Atoi(id)

	_, err = db.Exec(Queries.UpdateTask(idTask), task.Title, task.Description, task.CreatedAt, task.Deadline, task.Active)

	if err != nil {
		log.Fatalf("Invalid Id: %v", err)
	}

}
