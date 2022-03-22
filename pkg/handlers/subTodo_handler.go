package handlers

import (
	"fiber-apis/pkg/services"

	"github.com/gofiber/fiber/v2"
)

type SubTodoHandlerInterface interface {
	GetAllSubTodos(c *fiber.Ctx) error
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
