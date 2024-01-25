package utils

import (
	"fmt"
	"os"
)

// returns the value of the environment variable with the given key
func GetEnvValue(key string) (string, error) {

	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("%s environment variable not set", key)
	}

	return value, nil
}

// returns the database connection string
func GetDbConnectionString() (string, error) {

	host, err := GetEnvValue("DB_HOST")
	if err != nil {
		return "", err
	}

	port, err := GetEnvValue("DB_PORT")
	if err != nil {
		return "", err
	}

	user, err := GetEnvValue("DB_USER")
	if err != nil {
		return "", err
	}

	password, err := GetEnvValue("DB_PASSWORD")
	if err != nil {
		return "", err
	}

	dbname, err := GetEnvValue("DB_NAME")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname), nil
}
