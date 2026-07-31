package main

import "fmt"

func main() {
	intSlice := make([]int, 0, 5)
	intSlice = append(intSlice, 10, 12, 13, 143)
	fmt.Println(intSlice, len(intSlice), cap(intSlice))
}
