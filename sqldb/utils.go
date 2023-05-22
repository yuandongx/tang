package sqldb

import (
	"fmt"

	"tang/config"
)

type _error struct {
	e string
}

func (e _error) Error() string {
	return e.e
}
func error_(msg string) *_error {
	return &_error{e: msg}
}

func GetDbSource() string {
	host := config.GetEnv("DB_HOST")
	name := config.GetEnv("DB_NAME")
	username := config.GetEnv("DB_USERNAME")
	password := config.GetEnv("DB_PASSWORD")
	port := config.GetEnv("DB_PORT")
	if port == "" {
		port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		username, password, host, port, name)
}

func GetPostGresSource() string {
	host := config.GetEnv("DB_HOST")
	name := config.GetEnv("DB_NAME")
	username := config.GetEnv("DB_USERNAME")
	password := config.GetEnv("DB_PASSWORD")
	port := config.GetEnv("DB_PORT")
	if port == "" {
		port = "3306"
	}
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=verify-full",
		username, password, host, port, name)
}
