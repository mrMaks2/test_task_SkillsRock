package server

import (
	// "errors"
	// "fmt"
	"log"

	// "net/url"

	// "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	// _ "github.com/lib/pq"
	"test_task_SkillsRock/database"
	"test_task_SkillsRock/envs"
)

func InitServer() {

	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		log.Fatalf("Ошибка инициализации ENV: ", errEnvs)
	} else {
		log.Println("Инициализация ENV прошла успешно")
	}

	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatalf("Ошибка подключения к базе данных: ", errDatabase)
	} else {
		log.Println("Успешное подключение к базе данных")
	}

}

func StartServer() {
	// Инициализация роутеров
	InitRoutes()
}

// func ApplyMigrations() {

// 	env := envs.ServerEnvs
// 	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_HOST, env.POSTGRES_PORT, env.POSTGRES_NAME, env.POSTGRES_USE_SSL)

// 	// parsedURL, err := url.Parse(databaseURL)
// 	// if err != nil {
// 	// 	log.Fatalf("Ошибка парсинга URL: %v", err)
// 	// 	return
// 	// }

// 	m, err := migrate.New(
// 		"file://database/migrations",
// 		databaseURL)
// 	// parsedURL.String())
// 	if err != nil {
// 		log.Fatalf("Ошибка создания экземпляра миграции:", err)
// 	}

// 	err = m.Up()
// 	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
// 		log.Fatalf("Ошибка применения миграций: %v", err)
// 	}

// 	log.Println("Миграции успешно применены!")

// }
