package main

import (
	"fiber-apis/pkg/db"
	"fiber-apis/pkg/handlers"
	"fiber-apis/pkg/models"
	"fiber-apis/pkg/repositories"
	"fiber-apis/pkg/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database, err := db.Start()

	if err != nil {
		log.Fatal(err)
	}

	database.AutoMigrate(&models.Todo{}, &models.User{}, &models.SubTodo{})

	app := fiber.New()

	// Todo Injections
	todoRepo := repositories.NewTodoRepository(database)

	todoServices := services.TodoServiceStruct{
		Repository: todoRepo,
	}

	todoHandler := handlers.TodoHandlerStruct{
		Service: todoServices,
	}

	// Todo Routers
	app.Get("/todo/all", todoHandler.GetAllTodosHandler)

	app.Get("/todo/:id", todoHandler.GetTodoById)

	app.Post("/todo/create", todoHandler.CreateTodo)

	app.Delete("/todo/:id", todoHandler.DeleteTodoById)

	app.Put("/todo/update/:id", todoHandler.UpdateTodoById)

	// User Injections
	userRepo := repositories.UserRepoStruct{
		Db: database,
	}

	userService := services.UserServiceStruct{
		Repository: userRepo,
	}

	userHandler := handlers.UserHandlerStruct{
		Service: userService,
	}

	// User Routers
	app.Get("/user/all", userHandler.GetAllUsersHandler)

	app.Get("/user/:id", userHandler.GetUserById)

	app.Post("/user/create", userHandler.CreateUser)

	app.Delete("/user/delete/:id", userHandler.DeleteUserById)

	app.Put("/user/update/:id", userHandler.UpdateUserById)

	// SubTodo Injections
	subTodoRepo := repositories.SubTodoRepoStruct{
		Db: database,
	}
	subTodoService := services.SubTodoServiceStruct{
		Repository: subTodoRepo,
	}
	subTodoHandler := handlers.SubTodoHandlerStruct{
		Service: subTodoService,
	}

	// SubTodo Routers
	app.Get("/subTodo/all", subTodoHandler.GetAllSubTodos)

	app.Listen(":3000")
}
