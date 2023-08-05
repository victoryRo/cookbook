package chapter03

import (
	"fmt"
	"math"
)

type radian float64

func (rad radian) toDegrees() degree {
	return degree(float64(rad) * (180.0 / math.Pi))
}

// Float64 ...
func (rad radian) RFloat64() float64 {
	return float64(rad)
}

type degree float64

func (deg degree) toRadians() radian {
	return radian(float64(deg) * (math.Pi / 180.0))
}

// Float64 ...
func (deg degree) DFloat64() float64 {
	return float64(deg)
}

func DegreesAndRadians() {
	val := radiansToDegrees(1)
	fmt.Printf("One radian is : %.4f degrees\n", val)

	val2 := degreesToRadians(val)
	fmt.Printf("%.4f degrees is %.4f rad\n", val, val2)

	val = radian(1).toDegrees().DFloat64()
	fmt.Printf("Degrees: %.4f degrees\n", val)

	val = degree(val).toRadians().RFloat64()
	fmt.Printf("Rad: %.4f radians\n", val)
}

func degreesToRadians(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func radiansToDegrees(rad float64) float64 {
	return rad * (180.0 / math.Pi)
}

// La biblioteca estándar de Go no contiene ningún paquete con una función que convierta radianes en grados y viceversa.
// Pero al menos la constante Pi es parte del pkg math,
// por lo que la conversión podría realizarse como se muestra en el código de muestra.

// El código anterior también presenta el enfoque de definir el tipo personalizado con métodos adicionales.
// Estos simplifican la conversión de valores mediante una práctica API.
