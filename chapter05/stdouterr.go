package chapter05

import (
	"fmt"
	"io"
	"os"
)

// StdOutputError ...
// el destino donde se escriben los datos puede ser cualquier cosa,
// desde la consola hasta el socket.
// Esta receta le mostrará cómo escribir stdout y stderr.
func StdOutputError() {
	// Simplemente escribe una cadena
	io.WriteString(os.Stdout, "This is string to standard output.\n")

	io.WriteString(os.Stderr, "This is string to standard error output.\n")

	// Stdout/err implementa
	// interfaz de escritor
	buf := []byte{0xAF, 0xFF, 0xFe}
	for i := 0; i < 200; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}

	// El paquete fmt
	// también se puede usar
	fmt.Fprintln(os.Stdout)

	// Al igual que en Stdin la receta anterior,
	// Stdout y Stderr son los descriptores de archivo.
	// Así que estos están implementando la interfaz Writer
}
