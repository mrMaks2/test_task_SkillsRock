package main

import "test_task_SkillsRock/server"

func init() {
	server.InitServer()
	// server.ApplyMigrations()
}

func main() {
	server.StartServer()
}
