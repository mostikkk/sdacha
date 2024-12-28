package main

import (
	"log"
	"obuch/internal/database"
	"obuch/internal/handlers"
	"obuch/internal/taskService"
	"obuch/internal/userService"
	"obuch/internal/web/tasks"
	"obuch/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()

	if err := database.DB.AutoMigrate(&userService.Users{}); err != nil {

		log.Fatalf("failed to auto-migrate database: %v", err)

	}
	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {

		log.Fatalf("failed to auto-migrate database: %v", err)

	}

	taskRepo := taskService.NewTaskRepository(database.DB)
	TaskService := taskService.NewService(taskRepo)
	userRepo := userService.NewUsersRepository(database.DB)
	UserService := userService.NewService(userRepo)

	Taskhandler := handlers.NewTaskHandler(TaskService)
	Userhandler := handlers.NewUserHandler(UserService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTaskHandler := tasks.NewStrictHandler(Taskhandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)
	strictUserHandler := users.NewStrictHandler(Userhandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
