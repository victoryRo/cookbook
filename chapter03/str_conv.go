package chapter03

import (
	"fmt"
	"strconv"
)

// Conversión entre binario, octal, decimal y hexadecimal
// En algunos casos, los valores enteros pueden representarse mediante
// representaciones distintas de decimales.
// La conversión entre estas representaciones se realiza fácilmente
// con el uso del pkg strconv.

const binn = "10111"
const hexx = "1A"
const octt = "12"
const decc = "10"
const numd = "19"

// NumericalConversion ...
func NumericalConversion() {

	// Convierte el valor binario en hexadecimal
	v, _ := ConvertInt(binn, 2, 16)
	fmt.Printf("Binary value %s converted  to hex: %s\n", binn, v)

	// Convierte el valor hexadecimal en decimal
	v, _ = ConvertInt(hexx, 16, 10)
	fmt.Printf("Hex value %s converted to dec: %s\n", hexx, v)

	// Convierte el valor oct en hexadecimal
	v, _ = ConvertInt(octt, 8, 16)
	fmt.Printf("Oct value %s converted to hex: %s\n", octt, v)

	// Convierte el valor dec en oct
	v, _ = ConvertInt(decc, 10, 8)
	fmt.Printf("Dec value %s converted to oct: %s\n", decc, v)

	// Convierte el valor dec en binario
	v, _ = ConvertInt(numd, 10, 2)
	fmt.Printf("Dec value %s converted to %s\n", numd, v)
}

// ConvertInt convierte el valor de cadena dado de base
// para definir toBase.
func ConvertInt(val string, base, toBase int) (string, error) {
	// 1º param recibe el valor numerico como un string
	// 2º param base (2 binario) (8 octal) (hexa 16) etc
	// 3º param tamaño de bits (8) (32) (64)
	// retorna un int base 64
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}

	// 1º pparam recibe el valor int64
	// 2º param recibe la base como va ser representado base (2 binario) (8 octal) (16 hexa) etc
	// retorna un string con el numero representado en la base especificada
	return strconv.FormatInt(i, toBase), nil
}
