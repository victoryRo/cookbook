package chapter02

import (
	"bytes"
	"fmt"
)

// ConcantString ...
// convertimos un slice de string en un string
func ConcantString() {
	// tomamos un slices de string
	str := []string{"This ", "is ", "even ", "more ", "performant "}

	// inicializamos un buffer de bytes
	buf := bytes.Buffer{}

	// recorremos el slice str
	for _, val := range str {
		// agrega el contenido del val param al buffer
		// concatena el string
		buf.WriteString(val)
	}
	// retorna el contenido del buffer como un string
	fmt.Println(buf.String())
}

// ConcatCopy ...
func ConcatCopy() {
	// tomamos un slices de string
	str := []string{"This ", "is ", "even ", "more ", "performant "}

	// asigna e inicializa un obj de type
	bs := make([]byte, 100)
	bl := 0

	for _, val := range str {
		// copia de un slice a otro slice
		// copia de la fuente 2º param
		// al destino 1º param
		bl += copy(bs[bl:], []byte(val))
	}

	// toma los bytes[:] y los transforma a string()
	fmt.Println(string(bs[:]))
	fmt.Printf("Number of elements copied are: %d unds\n", bl)
}
