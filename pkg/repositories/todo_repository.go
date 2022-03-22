package repositories

import (
	"fiber-apis/pkg/models"

	"gorm.io/gorm"
)

type TodoRepositoryInterface interface {
	GetAllTodosRepository() ([]models.Todo, error)
	GetTodoById(id int) (models.Todo, error)
	CreateTodo(newTodo models.Todo) error
	DeleteTodoById(_id int) error
	UpdateTodoById(_id int, _todo models.Todo) error
}

type todoRepoStruct struct {
	db *gorm.DB
}

func NewTodoRepository(_db *gorm.DB) TodoRepositoryInterface {
	return todoRepoStruct{
		db: _db,
	}
}

func (r todoRepoStruct) GetAllTodosRepository() ([]models.Todo, error) {

	todos := []models.Todo{}

	return todos, r.db.Debug().Raw("SELECT * FROM `todos`").Scan(&todos).Error
}

func (r todoRepoStruct) GetTodoById(id int) (models.Todo, error) {

	todo := models.Todo{}

	return todo, r.db.Where("ID = ?", id).First(&todo).Error
}

func (r todoRepoStruct) CreateTodo(_newTodo models.Todo) error {

	newTodo := _newTodo

	return r.db.Create(&newTodo).Error
}

func (r todoRepoStruct) DeleteTodoById(_id int) error {

	return r.db.Where("ID = ?", _id).Delete(&models.Todo{}).Error
}

func (r todoRepoStruct) UpdateTodoById(_id int, _todo models.Todo) error {

	return r.db.Model(&models.Todo{}).Where("ID = ?", _id).Updates(&_todo).Error
}
