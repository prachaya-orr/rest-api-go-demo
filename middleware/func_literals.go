package main

import "fmt"

func lessThan(a, b int) bool {
	return a < b
}

func main() {
	//func literal
	r := func(a, b int) bool {
		return a < b
	}(2, 3)

	r1 := lessThan(2, 3)

	fmt.Println("result:", r)
	fmt.Println("result_1:", r1)
}
