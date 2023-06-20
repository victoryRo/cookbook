package samples

import (
	"flag"
	"fmt"
)

// UserFlags ...
func UserFlags() {
	fmt.Println("Thanks for you info Name City Gender and Age")

	name := flag.String("name", "", "user name")

	var city string
	flag.StringVar(&city, "city", "", "your city")

	var gender string
	flag.StringVar(&gender, "gender", "", "your gender")

	age := flag.Int("age", -1, "your age")

	flag.Parse()

	fmt.Printf("User data Name: %s \t City: %s \t Gender: %s \t Age: %d\n", *name, city, gender, *age)
}
