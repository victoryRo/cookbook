package chapter03

import (
	"fmt"
	"math"
)

var valA float64 = 3.55554444

// RoundNumber ...
func RoundNumber() {
	intVal := int(valA)
	fmt.Printf("Bad rounding %v of type %T by casting to int: %v\n", valA, valA, intVal)

	fRound := myRound(valA)
	fmt.Printf("Rounding by custom function: %v\n", fRound)

	// redondea un numero float a su valor mas cercano
	roundNum := math.Round(valA)
	fmt.Printf("Rounding by Round(): %v\n", roundNum)
}

func myRound(x float64) float64 {
	// trunca el float = 3
	t := math.Trunc(x)
	// resta t 3 a x 3.55554444
	// si el num abs es mayor a 5 se le agrega el 1
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}
