package chapter01

import (
	"fmt"
	"os"
)

// Parameterize ...
func Parameterize() {
	args := os.Args

	// Esta llamada se imprimirá
	// todos los argumentos de la línea de comandos.
	fmt.Println(args)

	// El primer argumento, elemento cero del segmento,
	// es el nombre del binario llamado.
	programName := args[0]
	fmt.Printf("The binary name is: %s \n", programName)

	// El resto de los argumentos se pudieron obtener
	// omitiendo el primer argumento.
	otherArguments := args[1:]
	fmt.Println(otherArguments)

	for idx, arg := range otherArguments {
		fmt.Printf("Arg %d = %s \n", idx, arg)
	}
}
