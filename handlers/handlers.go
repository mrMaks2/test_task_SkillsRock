package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"test_task_SkillsRock/database"
	"test_task_SkillsRock/models"
)

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	sql := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id`
	err := database.GetDB().QueryRow(context.Background(), sql, task.Title, task.Description, task.Status).Scan(&task.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func GetTasks(c *fiber.Ctx) error {
	rows, err := database.GetDB().Query(context.Background(), "SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	sql := `UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = now() WHERE id = $4`
	_, err := database.GetDB().Exec(context.Background(), sql, task.Title, task.Description, task.Status, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	sql := `DELETE FROM tasks WHERE id = $1`
	_, err := database.GetDB().Exec(context.Background(), sql, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
