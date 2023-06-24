package chapter01

import (
	"fmt"
	"os"
	"path/filepath"
)

// CurrentDirectory ...
func CurrentDirectory() {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Ruta al archivo ejecutable
	fmt.Println(exe)

	// Resolver el directorio
	// del ejecutable
	exePath := filepath.Dir(exe)
	fmt.Println("executable path: ", exePath)

	// Usa EvalSymlinks para obtener
	// el camino real.
	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Symlinks evaluated: " + realPath)
}
