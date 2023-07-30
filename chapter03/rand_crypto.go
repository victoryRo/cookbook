package chapter03

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
)

// RandNumber ...
func RandNumber() {
	// math/rand paquete, que es criptográficamente inseguro,
	// y nos permite generar la misma secuencia con el uso de Source con el mismo número de semilla.
	// Este enfoque se utiliza generalmente en las pruebas.
	// La razón para hacerlo es por la reproducibilidad de la secuencia.

	// retorna un nuevo numero al azar
	// unico param NewSource devuelve una nueva fuente pseudoaleatoria sembrada con el valor dado.
	sec1 := rand.New(rand.NewSource(10))
	sec2 := rand.New(rand.NewSource(10))

	for i := 0; i < 5; i++ {
		rnd1 := sec1.Int()
		rnd2 := sec2.Int()

		if rnd1 != rnd2 {
			fmt.Println("Rand generated non-equal sequence")
			break
		} else {
			fmt.Printf("Math/Rand1: %d , Math/Rand2: %d\n", rnd1, rnd2)
		}
	}

	for i := 0; i < 5; i++ {
		safeNum := newCryptoRand()
		safeNum2 := newCryptoRand()

		if safeNum == safeNum2 {
			fmt.Println("Crypto generated equal numbers")
			break
		} else {
			fmt.Printf("Crypto/Rand1: %d , Crypto/Rand2: %d\n", safeNum, safeNum2)
		}
	}
}

// crypto/rand paquete.
// La API utiliza para Reader proporcionar la instancia
// de un generador pseudoaleatorio criptográficamente fuerte.
// El paquete en sí tiene el valor predeterminado Reader que generalmente se basa
// en el generador de números aleatorios basado en el sistema.
func newCryptoRand() int64 {
	// devolver un valor aleatorio uniforme
	// 1º param Reader es una instancia global compartida de un generador de números aleatorios criptográficamente seguro.
	// 2º param NewInt asigna y devuelve un nuevo Int establecido en x.
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}

	return safeNum.Int64()
}
