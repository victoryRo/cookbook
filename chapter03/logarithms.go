package chapter03

import (
	"fmt"
	"math"
)

func Logarithm() {
	ln := math.Log(math.E)
	fmt.Printf("Ln(E) = %.4f\n", ln)

	log10 := math.Log10(-100)
	fmt.Printf("Log10(100) = %.4f\n", log10)

	log2 := math.Log2(2)
	fmt.Printf("Log2(2) = %.4f\n", log2)

	log_3_6 := Log(3, 6)
	fmt.Printf("Log3(6) = %.4f\n", log_3_6)
}

// Log calcula el logaritmo de
// base > 1 y x mayor 0
func Log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}

// Los logaritmos se utilizan en aplicaciones científicas,
// así como en visualizaciones y mediciones de datos.
// El paquete incorporado mathcontiene las bases del logaritmo comúnmente utilizadas.
// Usando estos, puedes obtener todas las bases.
