package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	res := add(0, 0)
	fmt.Print("0 + 0 = ")
	fmt.Println(res)
}
