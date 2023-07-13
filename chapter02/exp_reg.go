package chapter02

import (
	"fmt"
	"regexp"
)

//Siempre hay tareas como validar la entrada,
//buscar información en el documento o incluso limpiar una cadena determinada
//de caracteres de escape no deseados
//Para estos casos se suelen utilizar expresiones regulares.

const text = `[{\"email\": \"email@example.com\" \"phone\": 543777321}, 
              {\"email\": \"other@domain.com\" \"phone\": 098765333}]`

// ExpressionCheck ...
func ExpressionCheck() {
	// prepara la expresion regular
	emailRegexp := regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-z]+")

	// devuelve el string que contiene el texto de la primera coincidencia
	first := emailRegexp.FindString(text)
	fmt.Println("First: ")
	fmt.Println(first)

	// devuelve todas las coincidencias sucesivas de la expresión en []string
	all := emailRegexp.FindAllString(text, -1)
	fmt.Println("All: ")
	for _, val := range all {
		fmt.Println(val)
	}
}
