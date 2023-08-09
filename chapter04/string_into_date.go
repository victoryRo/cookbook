package chapter04

import (
	"fmt"
	"time"
)

// ParsingStringIntoDate ...
func ParsingStringIntoDate() {

	// Si la zona horaria no está definida
	// que devuelve la función Parse
	// la hora en la zona horaria UTC.
	t, err := time.Parse("2/1/2006", "31/7/2015")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// Tenga en cuenta que ParseInLocation
	// analiza el tiempo en una ubicación dada, si el
	// la cadena no contiene la definición de zona horaria
	t, err = time.ParseInLocation("2/1/2006 3:04 PM ", "31/7/2015 1:25 AM ", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
