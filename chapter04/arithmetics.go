package chapter04

import (
	"fmt"
	"time"
)

// DateArithmetics ...
// El tipo Time del paquete time también le permite realizar operaciones aritméticas
// básicas en la fecha y hora dadas.
// De esta manera, puede averiguar fechas pasadas y futuras.
func DateArithmetics() {
	l, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	t := time.Date(2017, 11, 30, 11, 10, 20, 0, l)
	fmt.Printf("Default day is: %v\n", t)

	// agregar 3 dias
	r1 := t.Add(72 * time.Hour)
	fmt.Printf("Default date +3HR is: %v\n", r1)

	// restar 3 dias
	r1 = t.Add(-72 * time.Hour)
	fmt.Printf("Default date -3HR is: %v\n", r1)

	// API más cómoda
	// para agregar días/meses/años
	r1 = t.AddDate(1, 3, 2)
	fmt.Printf("Default date +1YR +3MTH +2D is: %v\n", r1)
}
