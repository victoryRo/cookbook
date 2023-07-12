package chapter02

import (
	"fmt"
	"regexp"
	"strings"
)

// REEMPLAZA PARTE DEL STRING
// Otra tarea muy común relacionada con el procesamiento de cadenas
// es la sustitución de la subcadena en una cadena,
// La biblioteca estándar de Go proporciona la Replace función y el tipo Replacer
// para el reemplazo de varias cadenas a la vez.

const str = "Mary had a little lamb"
const str2 = "lamb lamb lamb lamb"

// ReplaceString ...
func ReplaceString() {
	out := strings.Replace(str, "lamb", "wolf", -1)
	fmt.Println(out)

	out = strings.Replace(str2, "lamb", "wolf", 2)
	fmt.Println(out)
}

//--------------------------------------------------------------------------------------------------

const str3 = "Mary had a little lamb"

// StringReplacer ...
func StringReplacer() {
	// organiza las palabras a remplazar por pares
	// de old a new
	replacer := strings.NewReplacer("lamb", "wolf", "Mary", "Jack")

	// pasamos el string y busca las conincidencias y las reemplaza
	out := replacer.Replace(str3)

	fmt.Println(out)
	// Jack had a little wolf
}

//--------------------------------------------------------------------------------------------------

const str4 = "Mary had a little lamb"

// RegexpReplace ...
func RegexpReplace() {

	regex := regexp.MustCompile("l[a-z]+")
	out := regex.ReplaceAllString(str4, "replacement")
	fmt.Println(out)
}
