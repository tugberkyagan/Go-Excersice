package services

import (
	"fiber-apis/pkg/models"
	"fiber-apis/pkg/repositories"
)

type SubTodoServiceInterface interface {
	GetAllSubTodos() ([]models.SubTodo, error)
}

type SubTodoServiceStruct struct {
	Repository repositories.SubTodoRepositoryInterface
}

func (s SubTodoServiceStruct) GetAllSubTodos() ([]models.SubTodo, error) {

	return s.Repository.GetAllSubTodos()

}
