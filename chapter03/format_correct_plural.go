package chapter03

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// PluralFormatting
// Al mostrar mensajes para el usuario, la interacción es más placentera
// si las oraciones se sienten más humanas. El paquete Go golang.org/x/text,
// que es el paquete de extensión, contiene esta característica para dar formato
// a los plurales de la manera correcta.
func PluralFormatting() {

	message.Set(language.English, "%d items to do",
		plural.Selectf(1, "%d",
			"=0", "no items to do",
			plural.One, "one item to do",
			"<100", "%[1]d items to do",
			plural.Other, "lot of items to do"))

	message.Set(language.English, "The average is %.2f",
		plural.Selectf(1, "%.2f",
			"<1", "The average is zero",
			"=1", "The average is one",
			plural.Other, "The average is %[1]f "))

	prt := message.NewPrinter(language.English)
	prt.Printf("%d items to do", 0)
	prt.Println()
	prt.Printf("%d items to do", 1)
	prt.Println()
	prt.Printf("%d items to do", 10)
	prt.Println()
	prt.Printf("%d items to do", 1000)
	prt.Println()

	prt.Printf("The average is %.2f", 0.8)
	prt.Println()
	prt.Printf("The average is %.2f", 1.0)
	prt.Println()
	prt.Printf("The average is %.2f", 10.0)
	prt.Println()
}

// SpanishForm ...
func SpanishForm() {
	message.Set(language.Spanish, "su ciudad es %d",
		plural.Selectf(1, "%d",
			"=10", "su ciudad es miami ok go",
			"<4", "su ciudad esta muy lejos",
			"<11", "su ciudad es muy bonita",
			plural.Other, "su ciudad no la conozco"))

	msg := message.NewPrinter(language.Spanish)
	msg.Printf("su ciudad es %d", 10)
	msg.Println()
	msg.Printf("su ciudad es %d", 2)
	msg.Println()
	msg.Printf("su ciudad es %d", 11)
	msg.Println()
	msg.Printf("su ciudad es %d", 9)
	msg.Println()
}
