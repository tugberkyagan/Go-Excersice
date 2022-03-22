package handlers

import (
	"fiber-apis/pkg/models"
	"fiber-apis/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandlerInterface interface {
	GetAllUsersHandler(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	DeleteUserById(c *fiber.Ctx) error
	UpdateUserById(c *fiber.Ctx) error
}

var _ UserHandlerInterface = UserHandlerStruct{}

type UserHandlerStruct struct {
	Service services.UserServiceInterface
}

func (h UserHandlerStruct) GetAllUsersHandler(c *fiber.Ctx) error {
	users, err := h.Service.GetAllUsersService()

	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	return c.JSON(users)
}

func (h UserHandlerStruct) GetUserById(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	response, err := h.Service.GetUserById(id)

	if err != nil {
		return err
	}

	return c.JSON(response)

}

func (h UserHandlerStruct) CreateUser(c *fiber.Ctx) error {
	user := &models.User{}

	if err := c.BodyParser(user); err != nil {
		return err
	}

	response := h.Service.CreateUser(*user)

	return response
}

func (h UserHandlerStruct) DeleteUserById(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	return h.Service.DeleteUserById(id)

}

func (h UserHandlerStruct) UpdateUserById(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	user := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	return h.Service.UpdateUserById(id, user)
}
