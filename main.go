package main

import "fmt"

func main() {
	fmt.Println(some(1, 1))
}

func some(a int, b int) int {
	return a + b
}
