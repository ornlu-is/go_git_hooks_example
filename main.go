package main

import "fmt"

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("1 is even?", isEven(1))
	fmt.Println("32 is even?", isEven(32))
}
