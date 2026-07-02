package main

import "fmt"

func main() {
	fmt.Println("----COPIA INSEGURA----")
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice1_copy = slice1

	fmt.Println("Slice original: ", slice1)
	fmt.Println("Slice copia: ", slice1_copy)

	slice1[0] = 10

	fmt.Println("Slice original: ", slice1)
	fmt.Println("Slice copia: ", slice1_copy)

	slice1_copy = append(slice1_copy, 6)
	slice1_copy[0] = 99

	fmt.Println("Slice original: ", slice1)
	fmt.Println("Slice copia: ", slice1_copy)

	fmt.Println("----FIM COPIA INSEGURA----")

	fmt.Println("----COPIA SEGURA----")
	var slice2 = []int{1, 2, 3, 4, 5}
	var slice2_copy = make([]int, len(slice2))
	copy(slice2_copy, slice2)

	fmt.Println("Slice original: ", slice2)
	fmt.Println("Slice copia: ", slice2_copy)

	slice2_copy[0] = 10

	fmt.Println("Slice original: ", slice2)
	fmt.Println("Slice copia: ", slice2_copy)
	fmt.Println("----FIM COPIA SEGURA----")

}
