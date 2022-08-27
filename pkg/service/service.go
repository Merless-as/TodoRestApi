package service

import (
	"TestRestApi"
	"TestRestApi/pkg/repository"
)

type Authorization interface {
	CreateUser(user TestRestApi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list TestRestApi.TodoList) (int, error)
	GetAll(userId int) ([]TestRestApi.TodoList, error)
	GetById(userId, listId int) (TestRestApi.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input TestRestApi.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item TestRestApi.TodoItem) (int, error)
	GetAll(userId, listId int) ([]TestRestApi.TodoItem, error)
	GetById(userId, itemId int) (TestRestApi.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input TestRestApi.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList: NewTodoListService(repos.TodoList),
		TodoItem: NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
