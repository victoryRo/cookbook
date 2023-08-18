package chapter05

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Lectura de entrada estándar
// Cada proceso posee su descriptor estándar de archivo de entrada, salida y error.
// El stdin sirve como entrada del proceso.
// Esta receta describe cómo leer los datos del archivo stdin.

// ReadingStandardInput ...
func ReadingStandardInput() {

	var name string
	fmt.Println("What is your name?")
	// escanea el texto leído desde la entrada estándar 'terminal'
	// toma los valores de tipo segun el formato
	fmt.Scanf("%s\n", &name)

	var age int
	fmt.Println("What is your age?")
	// escanea el texto leído desde la entrada estándar 'terminal'
	// toma los valores de tipo segun el formato
	fmt.Scanf("%d\n", &age)

	fmt.Printf("Hello %s your age is %d\n", strings.ToTitle(name), age)
}

// Scanner ...
func Scanner() {
	// El escáner es capaz de
	// escanear la entrada por líneas
	sc := bufio.NewScanner(os.Stdin)

	// Scan avanza el escáner al siguiente token
	// que estará disponible a través del método Bytes o Text.
	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Echo %s\n", txt)
	}
}

// Reader ...
func Reader() {

	for {
		// slice de bytes longitud 8 iniciados con valores 0
		data := make([]byte, 8)
		// desde terminal stdin lee y guarda en data hasta logitud de 8 bytes
		n, err := os.Stdin.Read(data)

		// si existen entradas muestra el resultado
		if err == nil && n > 0 {
			process(data)
		} else {
			break
		}
	}
}

func process(data []byte) {
	// muestra resultado como bytes y como string
	fmt.Printf("Received %X %s\n", data, string(data))
}

// echo 'Go is awesome!' | go run .
