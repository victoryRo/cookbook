package chapter02

import (
	"fmt"
	"regexp"
	"strings"
)

const phrase = "Mery had a little lamb"

// DivideWords ...
func DivideWords() {
	// toma un string y lo divide en palabras dentro de un slice
	words := strings.Fields(phrase)
	fmt.Printf("this value is: %#v and this type is: %T\n", words, words)

	// recorre el slice words e imprimimos las palabras
	for idx, word := range words {
		fmt.Printf("Word %d is %s\n", idx, word)
	}
}

const phrase2 = "Mary_had a little_lamb"

// DivideWords2 ...
func DivideWords2() {
	// toma un string y lo divide en palabras que contangan el simbolo
	// de separacion del 2º parametro
	words := strings.Split(phrase2, "_")
	fmt.Printf("this value is: %#v and this type is: %T\n", words, words)

	for i, w := range words {
		fmt.Printf("Word %d is %s\n", i, w)
	}
}

const phrase3 = "Mary*had,a%little_lamb"

// DivideWords3 ...
func DivideWords3() {
	// Se llama a splitFunc para cada
	// runa en un string. Si la runa
	// es igual a cualquier carácter en un "*%,_"
	// la phrase3 se divide.
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_", r)
	}

	// divide el string por c/u de las runas recoridas por el splitFunc()
	// que se encuentren el la phrase3 y retorna un slice
	words := strings.FieldsFunc(phrase3, splitFunc)

	for i, w := range words {
		fmt.Printf("Word %d is %s\n", i, w)
	}
}

const phrase4 = "Mary*had,a%little_lamb"
const phrase5 = "Hello_have_a=good=day=today"

// DivideRegex ...
func DivideRegex() {
	// especifica una expression regular para dividir la phrase4
	words := regexp.MustCompile("[*,%_]{1}").Split(phrase4, -1)

	for i, w := range words {
		fmt.Printf("Word %d is %s\n", i, w)
	}

	phrase := regexp.MustCompile("[=_]{1}").Split(phrase5, -1)
	for i, w := range phrase {
		fmt.Printf("Word %d is %s\n", i, w)
	}
}
