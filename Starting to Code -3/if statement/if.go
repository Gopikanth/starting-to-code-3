package main

import "fmt"

func main() {
	shopping := map[string]int{
		"computer": 100,
		"laptop":   200,
		"mobile":   300,
	}
	if _, ok := shopping["computer"]; ok {
		fmt.Println("Search Element Found")
	} else {
		fmt.Println("Search not Found")
	}
}
