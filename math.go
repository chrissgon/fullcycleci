package main

import "fmt"

func main() {
	fmt.Println(Some(1, 1))
}

func Some(a int, b int) int {
	return a + b
}
