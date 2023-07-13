package chapter01

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// ArrayValue Tipo personalizado necesario para implementar
// interfaz flag.Value para poder
// usarlo en la función flag.Var.
type ArrayValue []string

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

// Set ...
func (a *ArrayValue) Set(s string) error {
	*a = strings.Split(s, ",")
	return nil
}

// LearnFlag ...
func LearnFlag() {
	// Extrayendo valores de bandera con métodos que devuelven punteros
	retry := flag.Int("retry", -1, "Defines max retry count")

	// Leer la bandera usando la función XXXVar.
	// En este caso se debe definir la variable
	// antes de la bandera.
	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "Logger prefix")

	var arr ArrayValue
	flag.Var(&arr, "array", "Input array to iterate through.")

	// Ejecuta la función flag.Parse, para
	// lee las banderas de las variables definidas.
	// Sin esta llamada a la bandera
	// las variables permanecen vacías.
	flag.Parse()

	// Ejemplo de lógica no relacionada con banderas
	logger := log.New(os.Stdout, logPrefix, log.Ldate)

	retryCount := 0
	for retryCount < *retry {
		logger.Println("Retrying connection")
		logger.Printf("Sending array %v\n", arr)
		retryCount++
	}
}
