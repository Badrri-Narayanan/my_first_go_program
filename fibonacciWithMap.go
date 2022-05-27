package main

import (
	"fmt"
)

func GetFibonacci(size int) []int {
	if size < 1 {
		return []int{}
	}
	fibonacci := make(map[int]int)
	fibonacci[0] = 0
	fibonacci[1] = 1

	for i := 2; i <= size; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	result := make([]int, 0, len(fibonacci))
	for i := 0; i < len(fibonacci); i++ {
		result = append(result, fibonacci[i])
	}
	return result
}

func main() {
	result := GetFibonacci(5)

	fmt.Println(result)
}
