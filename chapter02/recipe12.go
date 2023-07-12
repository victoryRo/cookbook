package chapter02

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// IndentingAText ...
func IndentingAText() {
	text := "Hi! Go is awesome."
	text = indent(text, 10)
	fmt.Println(text)

	text = unindent(text, 3)
	fmt.Println(text)

	text = unindent(text, 10)
	fmt.Println(text)

	text = indentByRune(text, 4, '/')
	fmt.Println(text)
}

func indentByRune(input string, indent int, r rune) string {
	// retorna 1º param repetido las veces del 2º param
	return strings.Repeat(string(r), indent) + " " + input
}

func indent(input string, indent int) string {
	padding := indent + len(input)
	// este formato representa "%10s"
	// 10 numero de espacios antes del string
	return fmt.Sprintf("% "+strconv.Itoa(padding)+"s", input)
}

func unindent(input string, indent int) string {
	count := 0
	for _, val := range input {
		if unicode.IsSpace(val) {
			count++
		}
		if count == indent || !unicode.IsSpace(val) {
			break
		}
	}
	return input[count:]
}
