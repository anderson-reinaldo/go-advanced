package main

import "fmt"

func main() {
	matriz1 := [][]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println("----Shallow Copy----")

	matriz2 := make([][]int, len(matriz1))
	copy(matriz2, matriz1)

	matriz2[0][0] = 99

	fmt.Println("Matriz 1: ", matriz1)
	fmt.Println("Matriz 2: ", matriz2)
	fmt.Println("----Deep Copy----")
	matriz3 := [][]int{
		{1, 2},
		{3, 4},
	}
	matriz4 := deepCopy(matriz1)

	matriz4[0][0] = 99

	fmt.Println("Matriz 3: ", matriz3)
	fmt.Println("Matriz 4: ", matriz4)
}

func deepCopy(matriz [][]int) [][]int {
	dst := make([][]int, len(matriz))

	for i, slice := range matriz {
		dst[i] = make([]int, len(slice))
		copy(dst[i], slice)
	}

	return dst
}
