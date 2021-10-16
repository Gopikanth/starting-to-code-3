package main

import "fmt"

func main() {
	var i interface{} = 30
	switch i.(type) {
	case int:
		fmt.Println("this is integer")
	case string:
		fmt.Println("this is a string") //finding type of variable
	case float32:
		fmt.Println("this is a float")
	}
}
