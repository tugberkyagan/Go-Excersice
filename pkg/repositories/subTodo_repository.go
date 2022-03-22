package repositories

import (
	"fiber-apis/pkg/models"

	"gorm.io/gorm"
)

type SubTodoRepositoryInterface interface {
	GetAllSubTodos() ([]models.SubTodo, error)
}

type SubTodoRepoStruct struct {
	Db *gorm.DB
}

func (r SubTodoRepoStruct) GetAllSubTodos() ([]models.SubTodo, error) {

	subTodos := []models.SubTodo{}

	return subTodos, r.Db.Raw("SELECT * FROM sub_todos").Scan(&subTodos).Error
}
