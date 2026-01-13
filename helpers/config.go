package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

var Env = map[string]string{}

func SetUpConfig() {
	var err error
	Env, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("field to read env file", err)
	}

}

func GetEnv(key, val string) string {
	result := Env[key]
	if result == "" {
		result = val
	}

	return result
}
