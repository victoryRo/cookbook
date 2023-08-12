package chapter04

import (
	"fmt"
	"time"
)

// DifferenceBetweenTwoDates ...
// Encontrar la diferencia entre dos fechas no es una tarea inusual.
// Para esta operación, el paquete estándar Go time,
// respectivamente el tipo Time, proporciona métodos de apoyo
func DifferenceBetweenTwoDates() {

	l, err := time.LoadLocation("America/Bogota")
	if err != nil {
		panic(err)
	}

	t := time.Date(2000, 1, 1, 0, 0, 0, 0, l)
	t2 := time.Date(2000, 1, 3, 0, 0, 0, 0, l)
	fmt.Printf("First default date is %v\n", t)
	fmt.Printf("Second default date is %v\n", t2)

	// Sub() method conocer la diferencia entre dos fechas
	dur := t2.Sub(t)
	fmt.Printf("The duration between t and t2 is %v\n", dur)

	// devuelve el tiempo transcurrido desde t en nanosegundos
	dur = time.Since(t)
	fmt.Printf("The duration between now and t is %v\n", dur)

	// devuelve la duracion de tiempo hasta t en nanosegundos
	dur = time.Until(t)
	fmt.Printf("The duration between t and now is %v\n", dur)
}
