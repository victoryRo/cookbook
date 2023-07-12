package chapter02

import (
	"encoding/csv"
	"fmt"
	"os"
)

// CheckCommanSeparateValues ...
func CheckCommanSeparateValues() {
	// abre el archivo .csv
	file, err := os.Open("./chapter02/data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// crea un nuevo lector para archivos .csv
	reader := csv.NewReader(file)
	// numero de campos esperados por registro "linea"
	reader.FieldsPerRecord = 3
	reader.Comment = '#'

	for {
		// recorre los registro del archivo csv
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}

// CommanSeparateValues ...
func CommanSeparateValues() {
	file, err := os.Open("./chapter02/data_uncommon.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// crea nuevo lector para archivos csv
	reader := csv.NewReader(file)
	// especifica que tipo de coma separa los valores
	reader.Comma = ';'

	for {
		// recorre uno por uno los registros
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}
