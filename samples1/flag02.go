package samples1

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type ArrayValue []string

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func (a *ArrayValue) Set(s string) error {
	*a = strings.Split(s, ",")
	return nil
}

// SumMultiplierFlags ...
func SumMultiplierFlags() {
	fmt.Println("ingress four numbers for sum and multiplication")

	var arr ArrayValue
	flag.Var(&arr, "numbers", "input numbers in array to iterate")

	flag.Parse()
	fmt.Printf("array values %v\n", arr)

	var sum int
	var mul int

	for _, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("we have an error: %v", err)
		}
		sum += num
		mul = num * num
	}
	fmt.Printf("The sum of this number is: %d\t and multiply is: %d\n", sum, mul)
}
