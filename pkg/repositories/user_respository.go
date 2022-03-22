package repositories

import (
	"fiber-apis/pkg/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsersRepository() ([]models.User, error)
	GetUserById(_id int) (models.User, error)
	CreateUser(_newUser models.User) error
	DeleteUserById(_id int) error
	UpdateUserById(_id int, _user models.User) error
}

type UserRepoStruct struct {
	Db *gorm.DB
}

func (r UserRepoStruct) GetAllUsersRepository() ([]models.User, error) {

	users := []models.User{}

	return users, r.Db.Raw("SELECT * FROM users").Scan(&users).Error
}

func (r UserRepoStruct) GetUserById(_id int) (models.User, error) {

	user := models.User{}

	return user, r.Db.Where("ID = ?", _id).First(&user).Error
}

func (r UserRepoStruct) CreateUser(_newUser models.User) error {

	newUser := _newUser

	return r.Db.Create(&newUser).Error

}

func (r UserRepoStruct) DeleteUserById(_id int) error {

	return r.Db.Where("ID = ?", _id).Delete(&models.User{}).Error

}

func (r UserRepoStruct) UpdateUserById(_id int, _user models.User) error {

	return r.Db.Model(&models.User{}).Where("ID = ?", _id).Updates(&_user).Error

}
