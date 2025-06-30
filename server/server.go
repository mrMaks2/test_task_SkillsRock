package server

import (
	"os"
	"fmt"
	"log"
	"test_task_SkillsRock/database"
	"test_task_SkillsRock/envs"
)

func InitServer() {

	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		log.Fatal("Ошибка инициализации ENV: ", errEnvs)
	} else {
		log.Println("Инициализация ENV прошла успешно")
	}

	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("Ошибка подключения к базе данных: ", errDatabase)
	} else {
		log.Println("Успешное подключение к базе данных")
	}

}

func StartServer() {
	// Инициализация роутеров
	InitRoutes()
}

func applyMigrations() {

	env := envs.ServerEnvs
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_HOST, env.POSTGRES_PORT, env.POSTGRES_NAME, env.POSTGRES_USE_SSL)

	m, err := migrate.New(
		"file://database/migrations",
		databaseURL)
	if err != nil {
		log.Fatal("Ошибка создания экземпляра миграции:", err)
	}

	if err := m.Up(); err != nil &amp;&amp; err != migrate.ErrNoChange {
		log.Fatal("Ошибка применения миграций:", err)
	}

	log.Println("Миграции успешно применены!")

}