package chapter02

import (
	"fmt"
	"strings"
	"unicode"
)

const email = "Example@domain.com"
const name = "isaac newton"
const upc = "upc"
const i = "i"

const snakeCase = "first_name"

// TransformWords ...
func TransformWords() {
	// Para comparar la entrada del usuario
	// a veces es mejor
	// comparar la entrada en un mismo
	// caso.
	input := "Example@domain.com"
	input = strings.ToLower(input)
	emailToCompare := strings.ToLower(email)
	matches := input == emailToCompare
	fmt.Printf("Email matches: %t\n", matches)

	upcCode := strings.ToUpper(upc)
	fmt.Println("UPPER case: " + upcCode)

	// Este dígrafo tiene mayúsculas y minúsculas diferentes
	// titulo del caso.
	str := "ǳ"
	fmt.Printf("%s in upper: %s and title: %s\n", str, strings.ToUpper(str), strings.ToTitle(str))

	title := strings.ToTitle(i)
	titleTurk := strings.ToTitleSpecial(unicode.TurkishCase, i)
	if title != titleTurk {
		fmt.Printf("ToTitle is different: %#U vs. %#U\n", title[0], []rune(titleTurk)[0])
	}

	// En algunos casos la entrada
	// debe corregirse por si acaso.
	correctNameCase := strings.Title(name)
	fmt.Println("Corrected name: " + correctNameCase)

	firstNameCamel := toCamelCase(snakeCase)
	fmt.Println("Camel case: " + firstNameCamel)
	//firstName
}

func toCamelCase(input string) string {
	titleSpace := strings.Title(strings.Replace(input, "_", " ", -1))
	camel := strings.Replace(titleSpace, " ", "", -1)
	return strings.ToLower(camel[:1]) + camel[1:]
}
