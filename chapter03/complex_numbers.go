package chapter03

import (
	"fmt"
	"math/cmplx"
)

// Los números complejos generalmente se usan para aplicaciones
// y cálculos científicos. Go implementa números complejos como el tipo primitivo.
// Las operaciones específicas sobre números complejos son parte del pkg math/cmplx

// NumbersComplex ...
func NumbersComplex() {

	// los números complejos son
	// definido como real e imaginario
	// parte definida por float64
	a := complex(2, 3)

	fmt.Printf("Real part: %f \n", real(a))
	fmt.Printf("Complex part: %f \n", imag(a))

	b := complex(6, 4)

	// Todo común
	// los operadores son útiles
	c := a - b
	fmt.Printf("Difference : %v\n", c)
	c = a + b
	fmt.Printf("Sum :%v\n", c)
	c = a * b
	fmt.Printf("Product : %v\n", c)
	c = a / b
	fmt.Printf("Product : %v\n", c)

	conjugate := cmplx.Conj(a)
	fmt.Println("Complex number a's conjugate : ", conjugate)

	cos := cmplx.Cos(b)
	fmt.Println("Cosine of b : ", cos)
}
