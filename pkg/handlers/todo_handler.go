package handlers

import (
	"fiber-apis/pkg/models"
	"fiber-apis/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type TodoHandlerInterface interface {
	GetAllTodosHandler(c *fiber.Ctx) error
	GetTodoById(c *fiber.Ctx) error
	CreateTodo(c *fiber.Ctx) error
	DeleteTodoById(c *fiber.Ctx) error
	UpdateTodoById(c *fiber.Ctx) error
}

type TodoHandlerStruct struct {
	Service services.TodoServiceInterface
}

func (h TodoHandlerStruct) GetAllTodosHandler(c *fiber.Ctx) error {

	todos, err := h.Service.GetAllTodosService()

	if err != nil {
		return err
	}
	return c.JSON(todos)
}

func (h TodoHandlerStruct) GetTodoById(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	todo, err := h.Service.GetTodoById(id)

	if err != nil {
		return err
	}
	return c.JSON(todo)
}

func (h TodoHandlerStruct) CreateTodo(c *fiber.Ctx) error {

	todo := &models.Todo{}

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	response := h.Service.CreateTodo(*todo)

	return response
}

func (h TodoHandlerStruct) DeleteTodoById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	return h.Service.DeleteTodoById(id)
}

func (h TodoHandlerStruct) UpdateTodoById(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	todo := models.Todo{}

	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	return h.Service.UpdateTodoById(id, todo)
}
