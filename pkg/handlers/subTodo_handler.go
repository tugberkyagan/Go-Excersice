package handlers

import (
	"fiber-apis/pkg/models"
	"fiber-apis/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type SubTodoHandlerInterface interface {
	GetAllSubTodos(c *fiber.Ctx) error
	GetSubTodoById(c *fiber.Ctx) error
	CreateSubTodo(c *fiber.Ctx) error
	DeleteSubTodoById(c *fiber.Ctx) error
	UpdateSubTodoById(c *fiber.Ctx) error
}

type SubTodoHandlerStruct struct {
	Service services.SubTodoServiceInterface
}

func (h SubTodoHandlerStruct) GetAllSubTodos(c *fiber.Ctx) error {

	subTodos, err := h.Service.GetAllSubTodos()

	if err != nil {
		return err
	}
	return c.JSON(subTodos)

}

func (h SubTodoHandlerStruct) GetSubTodoById(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	subTodo, err := h.Service.GetSubTodoById(id)

	if err != nil {
		return err
	}

	return c.JSON(subTodo)
}

func (h SubTodoHandlerStruct) CreateSubTodo(c *fiber.Ctx) error {

	subTodo := &models.SubTodo{}

	if err := c.BodyParser(subTodo); err != nil {
		return err
	}

	response := h.Service.CreateSubTodo(*subTodo)

	return response
}

func (h SubTodoHandlerStruct) DeleteSubTodoById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	return h.Service.DeleteSubTodoById(id)
}

func (h SubTodoHandlerStruct) UpdateSubTodoById(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return err
	}

	subTodo := models.SubTodo{}

	if err := c.BodyParser(&subTodo); err != nil {
		return err
	}

	return h.Service.UpdateSubTodoById(id, subTodo)
}
