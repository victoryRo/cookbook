package chapter03

import (
	"fmt"
	"strconv"
)

const bin = "00001"
const hex = "2f"
const intString = "12"
const floatString = "12.3"

// StringToNumber ...
func StringToNumber() {
	// decimals
	// convierte un num string en un integer
	res, err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parse integer: %d\n", res)

	// Parsing hexadecimal
	// 1º recive un string hexadecimal
	// 2º base 16
	// 3º de 32 bits
	res64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parse hexadecimal: %d\n", res64)

	// Parsing binary values
	resBin, err := strconv.ParseInt(bin, 2, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parse binary: %d\n", resBin)

	// Parsing floating-points
	resFloat, err := strconv.ParseFloat(floatString, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parse float: %.3f\n", resFloat)
}
