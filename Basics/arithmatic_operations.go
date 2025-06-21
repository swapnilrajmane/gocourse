package Basics

import (
	"fmt"
	"math"
)

func main() {

	var a int = 15
	var b int = 11
	var result int

	result = a + b
	fmt.Println("Addition", result)

	result = a - b
	fmt.Println("subsctraction", result)

	result = a * b
	fmt.Println("multiplication", result)

	result = a / b
	fmt.Println("division", result)

	result = a % b
	fmt.Println("factorial", result)

	const p float64 = 56 / 15.0
	fmt.Println("floating", p)

	var maxint int64 = 3838383838383898677
	fmt.Println(maxint)
	maxint = maxint + 1
	fmt.Println(maxint)

	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}
