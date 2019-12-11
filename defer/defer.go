package main

import "fmt"

func main() {
	defer fmt.Println("Test1")
	fmt.Println("Test2")
}
