package main

import "fmt"

func main() {
	a := 10
	b := 20
	if a > b && b > 0 {
		fmt.Println("your value is true")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("error")
	}
}
