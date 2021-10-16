package main

import "fmt"

func main() {
	switch 2 {
	case 2:
		fmt.Println(1)
		fallthrough //prints whatever below that
	case 3:
		fmt.Println(2)
	case 4:
		fmt.Println(5)

	default:
		fmt.Println("")
	}
	switch 2 {
	case 2:
		fmt.Println(1)
		break //break the statement
	case 3:
		fmt.Println(2)
	case 4:
		fmt.Println(5)

	default:
		fmt.Println("")
	}
}
