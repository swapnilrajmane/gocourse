package main

import (
	"fmt"
	foo "net/http"
)

func main() {
	fmt.Println("hello, go")

	resp, err := foo.Get("https://jsonplaceholders.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return

	}

	defer resp.Body.Close()

	fmt.Println("HTTP Response Status:", resp.Status)
}
