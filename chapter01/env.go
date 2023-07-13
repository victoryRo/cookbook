package chapter01

import (
	"fmt"
	"log"
	"os"
)

// EnvVariables ...
func EnvVariables() {
	connStr := os.Getenv("DB_CONN")
	log.Printf("Connection string: %s", connStr)
}

// EnvNoSet ...
func EnvNoSet() {
	key := "DB_CONN"

	connStr, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("The env variable %s is not set.\n", connStr)
	}
	fmt.Println(connStr)
}

// EnvVarTest work correctly ok...
func EnvVarTest() {
	key := "DB_CONN"

	//set the environment variable
	err := os.Setenv(key, "postgres://as:as@example.com/pg?sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}
	val := getEnvDefault(key, "postgres://as:as@localhost/pg?sslmode=verify-full")
	log.Println("the value is:" + val)

	err = os.Unsetenv(key)
	if err != nil {
		log.Printf("Error the UnsetEnv function %v", err)
	}
	val = getEnvDefault(key, "postgres://as:as@127.0.0.1/pg?sslmode=verify-full")
	log.Println("the default value is:" + val)
}

func getEnvDefault(key, defVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defVal
	}
	return val
}
