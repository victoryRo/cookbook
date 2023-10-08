package chapter01

import (
	"fmt"
	"os"
	"path/filepath"
)

// CurrentDirectory ... study
// Recuperar el directorio de trabajo actual
// Otra fuente de información útil para la aplicación
// es el directorio donde se encuentra el binario del programa.
func CurrentDirectory() {
	// Ejecutable devuelve el nombre de la ruta del ejecutable que inició el proceso actual.
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Ruta al archivo ejecutable
	fmt.Println(exe)

	// Resolver el directorio del ejecutable
	exePath := filepath.Dir(exe)
	fmt.Println("executable path: ", exePath)

	// Usa EvalSymlinks para obtener el camino real.
	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Symlinks evaluated: " + realPath)

	// Construya el binario con el comando go build -o program

	// /Users/yosoy/go/src/Packt/cookbook/bin/currentDirectory
	// executable path:  /Users/yosoy/go/src/Packt/cookbook/bin
	// Symlinks evaluated: /Users/yosoy/go/src/Packt/cookbook/bin
}
