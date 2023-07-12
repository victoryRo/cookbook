package chapter02

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"

// FindString ...
// busqueda de coincidencias en las frases
func FindString() {
	lookFor := "lamb"
	// comprueba si el 2º parametro se encuentra en el 1º
	contain := strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	lookFor = "wolf"
	contain = strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	startWith := "Mary"
	// comprueba si el 2º parametro se encuentra al inicio del 1º
	starts := strings.HasPrefix(refString, startWith)
	fmt.Printf("The \"%s\" starts with \"%s\": %t \n", refString, startWith, starts)

	endWith := "lamb"
	// comprueba si el 2º parametro se encuentra al final del 1º
	ends := strings.HasSuffix(refString, endWith)
	fmt.Printf("The \"%s\" ends with \"%s\": %t \n", refString, endWith, ends)
}
