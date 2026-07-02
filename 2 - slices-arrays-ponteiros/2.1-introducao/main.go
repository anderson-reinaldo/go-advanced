package main

import (
	"fmt"
	"time"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	array := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Tamanho do Slice: ", len(slice))
	fmt.Println("Capacidade do Slice: ", cap(slice))

	fmt.Println("Tamanho do array: ", len(array))
	fmt.Println("Capacidade do array: ", cap(array))

	slice = append(slice, 6)
	fmt.Println("---------------")
	fmt.Println("Tamanho do Slice: ", len(slice))
	fmt.Println("Capacidade do Slice: ", cap(slice))

	printSlice(slice)
}

func printSlice(slice []int) {
	for _, value := range slice {
		fmt.Println(value)
		time.Sleep(1 * time.Second)
	}
}
