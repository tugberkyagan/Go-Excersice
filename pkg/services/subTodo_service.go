package services

import (
	"fiber-apis/pkg/models"
	"fiber-apis/pkg/repositories"
)

type SubTodoServiceInterface interface {
	GetAllSubTodos() ([]models.SubTodo, error)
	GetSubTodoById(_id int) (models.SubTodo, error)
	CreateSubTodo(_subTodo models.SubTodo) error
	DeleteSubTodoById(_id int) error
	UpdateSubTodoById(_id int, _subTodo models.SubTodo) error
}

type SubTodoServiceStruct struct {
	Repository repositories.SubTodoRepositoryInterface
}

func (s SubTodoServiceStruct) GetAllSubTodos() ([]models.SubTodo, error) {

	return s.Repository.GetAllSubTodos()

}

func (s SubTodoServiceStruct) GetSubTodoById(_id int) (models.SubTodo, error) {

	return s.Repository.GetSubTodoById(_id)

}

func (s SubTodoServiceStruct) CreateSubTodo(_subTodo models.SubTodo) error {

	return s.Repository.CreateSubTodo(_subTodo)

}

func (s SubTodoServiceStruct) DeleteSubTodoById(_id int) error {

	return s.Repository.DeleteSubTodoById(_id)
}

func (s SubTodoServiceStruct) UpdateSubTodoById(_id int, _subTodo models.SubTodo) error {

	return s.Repository.UpdateSubTodoById(_id, _subTodo)
}
