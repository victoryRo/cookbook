package chapter04

import (
	"fmt"
	"time"
)

// EpochTime ...
func EpochTime() {

	// Establecer la época de int64
	t := time.Unix(0, 0)
	fmt.Println(t)

	// consigue la época
	// Desde la instancia de tiempo
	epoch := t.Unix()
	fmt.Println(epoch)

	// Tiempo actual de época
	apochNow := time.Now().Unix()
	fmt.Printf("Epoch time in seconds: %d\n", apochNow)

	apochNano := time.Now().UnixNano()
	fmt.Printf("Epoch time in nano-seconds: %d\n", apochNano)
}

// La época es el sistema universal para describir el punto en el tiempo.
// El comienzo de la época se define como 00:00:00 1 Jan 1970 UTC.
// El valor de epoch es la cantidad de segundos desde la marca de tiempo,
// menos la cantidad de segundos bisiestos desde entonces.

// El paquete time y el tipo Time le brindan la capacidad de operar y averiguar
// el tiempo de época de UNIX.
