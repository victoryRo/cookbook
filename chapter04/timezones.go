package chapter04

import (
	"fmt"
	"time"
)

// ConvertingBetweenTimeZones ...
// Tratar con las zonas horarias es difícil.
// Una buena manera de manejar las diferentes zonas horarias es mantener una zona horaria
// como referencia en el sistema y convertir las otras si es necesario.
// Esta receta le muestra cómo se realiza la conversión de tiempo entre zonas horarias.
func ConvertingBetweenTimeZones() {
	// retorna la locacion con el nombre dado
	eur, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	// fecha con la zona horaria especificada
	t := time.Date(2000, 1, 1, 0, 0, 0, 0, eur)
	fmt.Printf("Original time: %v\n", t)

	// retorna la locacion con el nombre dado
	phx, err := time.LoadLocation("America/Phoenix")
	if err != nil {
		panic(err)
	}

	// In devuelve una copia de t que representa el mismo instante de tiempo,
	// pero con la información de ubicación de la copia establecida en loc para fines de visualización.
	t2 := t.In(phx)
	fmt.Printf("Converted time: %v\n", t2)
}
