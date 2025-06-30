package main

import "test_task_SkillsRock/server"

func init() {
	server.InitServer()
}

func main() {
	server.StartServer()
	server.applyMigrations()
}
