package chapter03

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var integer int64 = 32500
var floatNum float64 = 22000.456

// FormatNumbers ...
func FormatNumbers() {

	// impreme decimales
	fmt.Printf("%d \n", integer)

	// muestra el signo
	fmt.Printf("%+d \n", integer)

	// Imprimir en otra base X -16, o-8, b -2, d - 10
	fmt.Printf("%X \n", integer)
	fmt.Printf("%#X \n", integer)

	// Relleno con ceros a la izquierda
	fmt.Printf("%010d \n", integer)

	// Relleno izquierdo con espacios
	fmt.Printf("% 10d \n", integer)

	// Relleno a la derecha
	fmt.Printf("% -10d  %s\n", integer, "hello")

	// -------------------------------------------------------------

	// Numero de punto
	fmt.Printf("%f \n", floatNum)

	// con un limite de precision = 5
	fmt.Printf("%.5f \n", floatNum)

	// Con notacion cientifica
	fmt.Printf("%e \n", floatNum)

	// número de punto flotante
	// %e para exponentes grandes
	// o %f de lo contrario
	fmt.Printf("%g \n", floatNum)
}

const num = 100000.5678

// Localized ...
func Localized() {

	// muestra el formato de precio para EEUU
	p := message.NewPrinter(language.English)
	p.Printf(" %.2f \n", num)

	// muestra el formato de precio para Alemania
	p = message.NewPrinter(language.German)
	p.Printf(" %.2f \n", num)

	// muestra el formato de precio para latino america
	p = message.NewPrinter(language.Spanish)
	p.Printf(" %.2f \n", num)
}
