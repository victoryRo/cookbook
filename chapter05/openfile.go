package chapter05

import (
	"fmt"
	"io"
	"os"
)

// OpenMyFile ...
func OpenMyFile() {

	// abre un archivo para leerlo
	f, err := os.Open("files/file.txt")
	if err != nil {
		panic(err)
	}

	// lee desde un archivo y retorna un []byte
	c, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// string() convierte un []byte en un string
	fmt.Printf("### file content ###\n%s\n", string(c))
	// cerramos el archivo despues de leerlo
	f.Close()

	// ----------------------------------------------------------------

	// abrimos un archivo
	// si no existe lo creamos
	// si existe podemos read leer y write escribir
	// permisos de archivo rwx rwx rwx
	f, err = os.OpenFile("files/test.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// escribe sobre el archivo el texto sugerido
	io.WriteString(f, "Test string thanks")
	f.Close()
}
