package repositories

import (
	"fiber-apis/pkg/models"

	"gorm.io/gorm"
)

type SubTodoRepositoryInterface interface {
	GetAllSubTodos() ([]models.SubTodo, error)
	GetSubTodoById(_id int) (models.SubTodo, error)
	CreateSubTodo(_subTodo models.SubTodo) error
	DeleteSubTodoById(_id int) error
	UpdateSubTodoById(_id int, _subTodo models.SubTodo) error
}

type SubTodoRepoStruct struct {
	Db *gorm.DB
}

func (r SubTodoRepoStruct) GetAllSubTodos() ([]models.SubTodo, error) {

	subTodos := []models.SubTodo{}

	return subTodos, r.Db.Raw("SELECT * FROM sub_todos").Scan(&subTodos).Error
}

func (r SubTodoRepoStruct) GetSubTodoById(_id int) (models.SubTodo, error) {

	subTodo := models.SubTodo{}

	return subTodo, r.Db.Where("ID = ?", _id).First(&subTodo).Error
}

func (r SubTodoRepoStruct) CreateSubTodo(_subTodo models.SubTodo) error {

	newSubTodo := _subTodo

	return r.Db.Create(&newSubTodo).Error
}

func (r SubTodoRepoStruct) DeleteSubTodoById(_id int) error {

	return r.Db.Where("ID = ?", _id).Delete(&models.SubTodo{}).Error
}

func (r SubTodoRepoStruct) UpdateSubTodoById(_id int, _subTodo models.SubTodo) error {

	return r.Db.Model(&models.SubTodo{}).Where("ID = ?", _id).Updates(&_subTodo).Error
}
