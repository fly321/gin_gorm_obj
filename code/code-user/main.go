package main

import "fmt"

func main() {
	var a int
	var b int
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		return
	}
	fmt.Println(a + b)
}
