package main

import (
	"fmt"
)

func main() {
	input := 0

	fmt.Scan(&input)
	num := 0
	for i := 0; i < 32; i++ {
		if input&1 == 1 {
			num++
		}
		input = input >> 1
	}
	fmt.Print(num)

}
