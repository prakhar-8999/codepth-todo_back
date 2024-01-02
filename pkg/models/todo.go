package models

import (
	"github.com/jinzhu/gorm"
	"github.com/prakhar-8999/todo/pkg/config"
)

var db *gorm.DB

type Todo struct {
	gorm.Model
	Task       string `gorm:""json:"task"`
	TaskStatus string `json:"taskStatus"`
	Status     string `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
}

func (t *Todo) CreateTodos() *Todo {
	db.NewRecord(t)
	db.Create(&t)
	return t
}

func GetTodos() []Todo {
	var Todos []Todo
	db.Find(&Todos)
	return Todos
}

func GetTodoById(ID int64) (*Todo, *gorm.DB) {
	var getTodo Todo
	db := db.Where("ID=?", ID).Find(&getTodo)
	return &getTodo, db
}

func DeleteTodo(ID int64) Todo {
	var todo Todo
	db.Where("ID=?", ID).Delete(todo)
	return todo
}
