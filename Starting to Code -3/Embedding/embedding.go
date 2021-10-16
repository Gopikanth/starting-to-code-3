package main

import "fmt"

type processor struct {
	processor_name string
	cores          int
}
type memory struct {
	memory_type     string
	memory_capacity int
}
type computer struct {
	processor
	memory
	price int
}

func main() {
	Computer := computer{}
	Computer.price = 5000
	Computer.processor_name = "AMD RYZON" // one way of declaration
	Computer.cores = 9
	Computer.memory_type = "DDR7"
	Computer.memory_capacity = 100
	computer_1 := computer{
		processor: processor{
			processor_name: "intel",
			cores:          8,
		},
		memory: memory{
			memory_type:     "dual_core", // other way of declaration
			memory_capacity: 11,
		},
		price: 1234,
	}

	fmt.Println(Computer)
	fmt.Println(computer_1)
	fmt.Println(computer_1.cores) //accessing a particular value in the struct

}
