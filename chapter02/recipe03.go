package chapter02

import (
	"fmt"
	"strings"
)

const selectBase = "SELECT * FROM user WHERE %s"

var refStringSlice = []string{
	" FIRST_NAME = 'Jack' ",
	" INSURANCE_NO = 333444555 ",
	" EFFECTIVE_FROM = SYSDATE ",
}

// JoinString ...
func JoinString() {
	// toma un slice de string 1º param y lo concatena
	// con el texto o simbol del 2º param
	sentence := strings.Join(refStringSlice, "AND")
	fmt.Printf(selectBase+"\n", sentence)
	// SELECT * FROM user WHERE FIRST_NAME = 'Jack' AND INSURANCE_NO = 333444555 AND EFFECTIVE_FROM = SYSDATE
}

//---------------------------------------------------------------------------------------------------------------

const selectText = "SELECT * FROM user WHERE "

type joinFunc func(piece string) string

// JoinString2 ...
func JoinString2() {
	jF := func(p string) string {
		if strings.Contains(p, "INSURANCE") {
			return "OR"
		}

		return "AND"
	}

	result := joinWithFunc(refStringSlice, jF)
	fmt.Println(selectText + result)
	// SELECT * FROM user WHERE  FIRST_NAME = 'Jack' OR INSURANCE_NO = 333444555 AND EFFECTIVE_FROM = SYSDATE
}

// joinWithFunc retorna un string
func joinWithFunc(stringSlice []string, joinFunc joinFunc) string {
	concatenate := stringSlice[0]

	// recorre el slice a partir de la posicion 1
	for _, val := range stringSlice[1:] {
		//             first text    OR AND        valor de string loop
		concatenate = concatenate + joinFunc(val) + val
	}
	return concatenate
}
