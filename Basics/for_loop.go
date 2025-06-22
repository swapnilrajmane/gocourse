package main

import "fmt"

func main() {
	// for i := 1; i <= 5; i++ {

	// 	fmt.Println(i)
	// }
	// iterative collection
	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for index, value := range numbers {
	//	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	// for i := 1; i <= 1000; i++ {
	//	if i%2 != 0 {
	//	continue
	//	}
	//	fmt.Println("even number is: ", i)
	//	if i == 899 {
	//		break

	//}

	//}

	// for i := range 10 {
	//	i++
	//	fmt.Println(10 - i)
	//}
	//fmt.Println("We have lift off!")

	rows := 10

	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 4*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
