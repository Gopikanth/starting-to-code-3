package main

import "fmt"

func main() {
	switch 2 {
	case 1:
		fmt.Println("this is case 1")
	case 2:
		fmt.Println("this is case 2") //single case
	case 3:
		fmt.Println("this is case three")
	default:
		fmt.Println("this is default")
	}
	switch 3 {
	case 1, 3, 5, 7:
		fmt.Println("this is odd")
	case 2, 4, 6, 8: //several cases
		fmt.Println("this is even")
	}
	switch a := 2 + 2; a {
	case 1, 3, 5, 7: //based on some conditions
		fmt.Println("this is odd")
	case 2, 4, 6, 8:
		fmt.Println("this is even")
	}
	i := 10
	switch {
	case i < 0:
		fmt.Println("negative Number") //switching like if_else block
	case i > 0:
		fmt.Println("Positive Number")

	}
}
