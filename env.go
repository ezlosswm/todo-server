package main

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (*DBConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &DBConfig{
		host:     os.Getenv("HOSTADDR"),
		user:     os.Getenv("USER_NAME"),
		dbname:   os.Getenv("DB_NAME"),
		password: os.Getenv("PASSWORD"),
	}, nil
}
