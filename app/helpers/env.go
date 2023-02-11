package helpers

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}

func ContextTimeOut() (timeOut int) {
	val := LoadEnv("CONTEXT_TIME_OUT")
	num, err := strconv.Atoi(val)
	if err != nil {
		return 10
	}
	return num

}
