package server

import (
	"github.com/gofiber/fiber/v2"
	"test_task_SkillsRock/envs"
	"test_task_SkillsRock/handlers"
)

func InitRoutes() {
	app := fiber.New()

	app.Post("/tasks", handlers.CreateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)
	app.Get("/tasks", handlers.GetTasks)
	app.Put("/tasks/:id", handlers.UpdateTask)

	app.Listen(":" + envs.ServerEnvs.CONN_HOST)
}
