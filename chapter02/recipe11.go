package chapter02

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// CleanupWhiteSpace ...
func CleanupWhiteSpace() {
	stringToTrim := "\t\t\n Go \tis\t Awesome \t\t"
	// limpia los espacios en blanco iniciales y finales en un string
	trimResult := strings.TrimSpace(stringToTrim)
	fmt.Println(trimResult)

	stringWithSpaces := "\t\t\n Go \tis\n Awesome \t\t"
	// el código representa el uso de la expresión regular
	// para reemplazar todos los múltiples espacios en blanco con un solo espacio:
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(stringWithSpaces, " ")
	fmt.Println(replace)

	needSpace := "need space"
	fmt.Println(pad(needSpace, 14, "CENTER"))
	fmt.Println(pad(needSpace, 14, "LEFT"))

	str := cleanAllBlankspaces(stringToTrim)
	fmt.Println(str)
	str2 := cleanAllBlankspaces(stringWithSpaces)
	fmt.Println(str2)
}

func cleanAllBlankspaces(textToClean string) string {
	// limpia los espacios interiores
	rex := regexp.MustCompile("\\s+")
	str := rex.ReplaceAllString(textToClean, " ")
	// limpia los espacios exteriores init final
	clean := strings.TrimSpace(str)

	return clean
}

func pad(input string, padLen int, align string) string {
	inputLen := len(input)

	if inputLen >= padLen {
		return input
	}

	repeat := padLen - inputLen
	var output string

	switch align {
	case "RIGHT":
		output = fmt.Sprintf("% "+strconv.Itoa(padLen)+"s", input)
	case "LEFT":
		output = fmt.Sprintf("% "+strconv.Itoa(padLen)+"s", input)
	case "CENTER":
		bothRepeat := float64(repeat) / float64(2)
		left := int(math.Floor(bothRepeat)) + inputLen
		right := int(math.Ceil(bothRepeat))
		output = fmt.Sprintf("% "+strconv.Itoa(left)+"s%"+strconv.Itoa(right)+"s", input, "")
	}

	return output
}
