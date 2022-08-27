package repository

import (
	"TestRestApi"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user TestRestApi.User) (int, error)
	GetUser(username, password string) (TestRestApi.User, error)
}

type TodoList interface {
	Create(userId int, list TestRestApi.TodoList) (int, error)
	GetAll(userId int) ([]TestRestApi.TodoList, error)
	GetById(userId, listId int) (TestRestApi.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input TestRestApi.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item TestRestApi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]TestRestApi.TodoItem, error)
	GetById(userId, itemId int) (TestRestApi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input TestRestApi.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}
