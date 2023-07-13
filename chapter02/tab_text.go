package chapter02

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// En ciertos casos, la salida (generalmente la salida de datos)
// se realiza a través de texto tabulado, que se formatea en celdas bien organizadas.
//	Este formato podría lograrse con el text/tabwriter pkg.
//	El paquete proporciona el Writer filtro,
// que transforma el texto con los caracteres de tabulación
// en texto de salida con el formato correcto.

// TabText ...
func TabText() {
	// la función crea el filtro Writer con los parámetros configurados.
	w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ', tabwriter.AlignRight)

	fmt.Fprintln(w, "username\tfirstname\tlastname\t")
	fmt.Fprintln(w, "sohlich\tRadomir\tSohlich\t")
	fmt.Fprintln(w, "novak\tJohn\tSmith\t")

	// debe ser llamado despues de la ultima llamada al writer
	w.Flush()
}

// TabTextExample ...
func TabTextExample() {
	// creamos el formato o diseño de tab o espacios para el texto que usaremos
	write := tabwriter.NewWriter(os.Stdout, 20, 0, 3, ' ', 2)

	// pasamos el formato tipo io.writer y el texto a ser formateado
	fmt.Fprintln(write, "Study\tTeacher\tResult\t")
	fmt.Fprintln(write, "Victor\tMaster\t10/10\t")
	fmt.Fprintln(write, "Angela\tSomeone\t08/10\t")
	fmt.Fprintln(write, "Marisa\tExpert\t09/10\t")
	fmt.Fprintln(write, "Paulo\tLife\t08/10\t")

	write.Flush()
}
