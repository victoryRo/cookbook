package chapter04

import (
	"fmt"
	"time"
)

// CurrentTime ...
func CurrentTime() {
	today := time.Now()
	fmt.Println(today)
	// año mes dia hora minitos segundos
	// 2023-08-07 08:47:09.158439 -0500 -05 m=+0.000112433
}

// TimeString ...
func TimeString() {
	tTime := time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)

	// El formateo está hecho
	// con uso de valor de referencia
	// 2 de enero 15:04:05 2006 MST
	// pasamos la forma como string y representa nuestra fecha en la forma dada
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/2"))
	// tTime is: 2017/3/5

	// pasamos la forma como string y representa nuestra hora en la forma dada
	fmt.Printf("The time is: %s\n", tTime.Format("15:04"))
	// The time is: 08:05

	// Los formatos predefinidos pueden usarse
	fmt.Printf("The time is: %s\n", tTime.Format(time.RFC1123))
	// The time is: Sun, 05 Mar 2017 08:05:02 -05

	// El formato admite espacio de relleno
	// solo por días en Go versión 1.9.2
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/_2"))
	// tTime is: 2017/3/ 5

	// El relleno de ceros se hace agregando 0
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/01/02"))
	// tTime is: 2017/03/05

	// La fracción con ceros a la izquierda usa 0s
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.00"))
	// tTime is: 08:05:02.00

	// La fracción sin ceros a la izquierda usa 9s
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.999"))
	// tTime is: 08:05:02

	// Agregar formato agrega el tiempo formateado a dado
	// buffer
	fmt.Println(string(tTime.AppendFormat([]byte("The time is up: "), "03:04PM")))
	// The time is up: 08:05AM
}
