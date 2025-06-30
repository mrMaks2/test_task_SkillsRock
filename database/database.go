package database

import (
	"context"
	"fmt"
	"log"
	"test_task_SkillsRock/envs"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var DB *pgxpool.Pool

func InitDatabase() error {
	var err1, err2 error
	env := envs.ServerEnvs
	uri := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_HOST, env.POSTGRES_PORT, env.POSTGRES_NAME, env.POSTGRES_USE_SSL)
	config, err1 := pgxpool.ParseConfig(uri)
	if err1 != nil {
		log.Fatalf("Не удалось распарсить строку подключения: %v\n", err1)
	}
	db, err2 := pgxpool.NewWithConfig(context.Background(), config)
	if err2 != nil {
		return err2
	} else {
		DB = db
		return nil
	}
}

func GetDB() *pgxpool.Pool {
	return DB
}
