package chapter04

import (
	"fmt"
	"time"
)

// RetrievingTimeUnits ...
// El tipo Time también proporciona la API para recuperar unidades de tiempo de la instancia.
// Esto significa que puede averiguar qué día del mes o qué hora del día representa la instancia.
// Esta receta muestra cómo obtener dichas unidades.
func RetrievingTimeUnits() {
	// definimos una fecha explicita con Date method
	t := time.Date(2017, 11, 29, 21, 0, 0, 0, time.Local)
	fmt.Printf("Extracting units from: %v\n", t)

	// podemos acceder a la instancia y recuperar diferentes valores
	dOfMonth := t.Day()    // dia del mes
	weekDay := t.Weekday() // dia de la semana
	month := t.Month()     // mes

	fmt.Printf("The %dth day of %v is %v\n", dOfMonth, month, weekDay)

	t2 := time.Now()
	fmt.Printf("Extracting units from: %v\n", t2)

	dOfMonth2 := t2.Day()
	month2 := t2.Month()
	weekDay2 := t2.Weekday()
	year := t2.Year()

	fmt.Printf("The %dth day of %v is %v in year %v\n", dOfMonth2, month2, weekDay2, year)
}
