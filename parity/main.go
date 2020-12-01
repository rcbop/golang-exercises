package main

import "fmt"

func main() {
	integerSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range integerSlice {
		if number%2 == 0 {
			fmt.Printf("%d is even\n", number)
		} else {
			fmt.Printf("%d is odd\n", number)
		}
	}
}
