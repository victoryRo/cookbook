package samples1

import (
	"fmt"
	"log"
	"os"
)

// EnvConfig ...
func EnvConfig() {
	//configuro una variable de entorno
	key := "DB_NAME"
	key2 := ""
	err := os.Setenv(key, "SecurityDB")
	if err != nil {
		log.Printf("Env not config err: %v\n", err)
	}

	val := getEnvConfig(key)
	fmt.Printf("The value of env is: %s\n", val)

	val = getEnvConfig(key2)
	fmt.Printf("the value of env is not set: %s\n", val)

}

// getEnvConfig comprueba si la env esta configurada
// con un valor valido
func getEnvConfig(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("Env not found key empty: %v", key)
	}
	return val
}
