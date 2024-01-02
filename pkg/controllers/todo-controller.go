package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prakhar-8999/todo/pkg/models"
	"github.com/prakhar-8999/todo/pkg/utils"
)

var NewTodo models.Todo

func GetTodos(w http.ResponseWriter, r *http.Request) {
	newTodo := models.GetTodos()
	res, _ := json.Marshal(newTodo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTodos(w http.ResponseWriter, r *http.Request) {
	CreateTodo := &models.Todo{}
	utils.ParseBody(r, CreateTodo)
	b := CreateTodo.CreateTodos()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todo_id := vars["todo_id"]
	ID, err := strconv.ParseInt(todo_id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todo := models.DeleteTodo(ID)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var updateTodo = &models.Todo{}
	utils.ParseBody(r, updateTodo)
	vars := mux.Vars(r)
	todo_id := vars["todo_id"]
	ID, err := strconv.ParseInt(todo_id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todoDetails, db := models.GetTodoById(ID)

	if updateTodo.Task != "" {
		todoDetails.Task = updateTodo.Task
	}
	if updateTodo.Status != "" {
		todoDetails.Status = updateTodo.Status
	}
	if updateTodo.TaskStatus != "" {
		todoDetails.TaskStatus = updateTodo.TaskStatus
	}
	db.Save(&todoDetails)
	res, _ := json.Marshal(todoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
