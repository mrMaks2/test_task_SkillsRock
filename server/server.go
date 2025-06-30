package server

import (
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
