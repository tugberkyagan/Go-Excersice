package services

import (
	"fiber-apis/pkg/models"
	"fiber-apis/pkg/repositories"
)

type TodoServiceInterface interface {
	GetAllTodosService() ([]models.Todo, error)
	GetTodoById(id int) (models.Todo, error)
	CreateTodo(newTodo models.Todo) error
	DeleteTodoById(_id int) error
	UpdateTodoById(_id int, _todo models.Todo) error
}

type TodoServiceStruct struct {
	Repository repositories.TodoRepositoryInterface
}

func (s TodoServiceStruct) GetAllTodosService() ([]models.Todo, error) {

	todo, err := s.Repository.GetAllTodosRepository()

	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s TodoServiceStruct) GetTodoById(id int) (models.Todo, error) {

	todo, err := s.Repository.GetTodoById(id)

	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (s TodoServiceStruct) CreateTodo(_newTodo models.Todo) error {

	newTodo := s.Repository.CreateTodo(_newTodo)

	return newTodo
}

func (s TodoServiceStruct) DeleteTodoById(_id int) error {

	response := s.Repository.DeleteTodoById(_id)

	return response
}

func (s TodoServiceStruct) UpdateTodoById(_id int, _todo models.Todo) error {

	return s.Repository.UpdateTodoById(_id, _todo)

}
