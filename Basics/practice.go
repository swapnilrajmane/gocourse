package main

import (
	"fmt"
)

var i int

func main() {

	a := 5
	b := 10
	var result int

	result = a + b
	fmt.Println("addtion", result)

	result = a - b
	fmt.Println("Substraction", result)

	result = a * b
	fmt.Println("multiplication", result)

	result = a / b
	fmt.Println("division", result)

	result = a % b
	fmt.Println("factorial", result)

	const p float64 = 10 / 3.0
	fmt.Println(p)

	for i = 1; i <= 10; i++ {
		fmt.Println(10 - i)

	}
	fmt.Println("lift off")

	for i = 1; i <= 100; i++ {

		if i%2 != 0 {
			continue
		}
		fmt.Println("this is even number:", i)
		if i == 34 {
			break
		}

	}

}
