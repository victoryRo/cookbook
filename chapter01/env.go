package chapter01

import (
	"fmt"
	"log"
	"os"
)

// EnvVariables ... study
func EnvVariables() {
	// recupera el valor de la env
	// pero no comprueba si esta vacio o no ha sido establecida
	connStr := os.Getenv("DB_CONN")
	l := len(connStr)
	log.Printf("Connection string: %s len: %d", connStr, l)
}

// EnvNoSet ...
func EnvNoSet() {
	key := "DB_CONN"
	// os.Setenv(key, "Whife")

	// recupera el valor de la env, si esta vacia retorna false
	connStr, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("The env variable %s is not set.\n", connStr)
	}
	fmt.Println(connStr)
}

// EnvVarTest work correctly ok...
func EnvVarTest() {
	key := "DB_CONN"

	// le asignamos el valor a la env
	err := os.Setenv(key, "first_environment_variable")
	if err != nil {
		log.Fatal(err)
	}
	val := getEnvDefault(key, "send_var_by_default")
	log.Println("the value is: " + val)

	// desabilita un variable de entorno
	err = os.Unsetenv(key)
	if err != nil {
		log.Printf("Error the UnsetEnv function %v", err)
	}
	val = getEnvDefault(key, "second_var_by_default")
	log.Println("the default value is :" + val)
}

// getEnvDefault ...
// comprueba si exite la env y la retorna, si no le asignamos una por default
func getEnvDefault(key, defVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defVal
	}
	return val
}
