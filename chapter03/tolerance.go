package chapter03

import (
	"fmt"
	"math"
	"math/big"
)

const da = 0.29999999999999998889776975374843459576368331909180
const db = 0.3

// NumberTolerance ...
func NumberTolerance() {
	daStr := fmt.Sprintf("%.10f", da)
	dbStr := fmt.Sprintf("%.10f", db)

	fmt.Printf("Strings %s = %s equals: %v \n", daStr, dbStr, dbStr == daStr)
	fmt.Printf("Number equals: %v \n", db == da)

	// Como la precisión de la representación flotante
	// está limitado. Para la comparación de flotación es
	// mejor usar la comparación con cierta tolerancia.
	fmt.Printf("Number equals with TOLERANCE: %v \n", equals(da, db))
}

// TOLERANCE ...
const TOLERANCE = 1e-8

// Equals compara los números de punto flotante
// con tolerancia 1e-8
func equals(numA, numB float64) bool {
	// retorna el numero absoluto de la resta
	delta := math.Abs(numA - numB)
	if delta < TOLERANCE {
		return true
	}
	return false
}

// -------------------------------------------------------------------------------------------

var de float64 = 0.299999992
var dd float64 = 0.299999991

var prec uint = 32
var prec2 uint = 16

// BigNumber ...
func BigNumber() {
	fmt.Printf("Comparing float64 with '==' equals: %v\n", de == dd)

	// genera un nuevo float y establece la nueva precision
	deB := big.NewFloat(de).SetPrec(prec)
	ddB := big.NewFloat(dd).SetPrec(prec)

	fmt.Printf("A: %v type: %T\n", deB, deB)
	fmt.Printf("B: %v type: %T\n", ddB, ddB)
	// hace una comparacion de los dos valores *big.Float
	fmt.Printf("Comparin big.Float with precision: %d : %v\n", prec, deB.Cmp(ddB) == 0)

	// genera un nuevo float y establece la nueva precision
	deB = big.NewFloat(de).SetPrec(prec2)
	ddB = big.NewFloat(dd).SetPrec(prec2)

	fmt.Printf("A: %v type %T \n", deB, deB)
	fmt.Printf("B: %v type %T \n", ddB, ddB)
	// hace una comparacion de los dos valores *big.Float
	fmt.Printf("Comparing big.Float with precision: %d : %v\n", prec2, deB.Cmp(ddB) == 0)
}
