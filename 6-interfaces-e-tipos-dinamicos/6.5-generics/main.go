package main

import "fmt"

func SomarInteiros(numeros []int) int {
	var soma int = 0

	for i := 0; i < len(numeros); i++ {
		soma += numeros[i]
	}

	return soma
}

func SomarFloat(numeros []float64) float64 {
	var soma float64 = 0

	for i := 0; i < len(numeros); i++ {
		soma += numeros[i]
	}

	return soma
}

func SomarAny(numeros []any) any {
	var soma float64 = 0

	for i := 0; i < len(numeros); i++ {
		switch v := numeros[i].(type) {
		case float64:
			soma += v
		case int:
			soma += float64(v)
		}
	}

	return soma
}

func SomaGenerics[T int | float64](numeros []T) T {
	var soma T = 0

	for i := 0; i < len(numeros); i++ {
		soma += numeros[i]
	}

	return soma
}

func main() {
	var (
		listaInteiros = []int{1, 2, 3, 4, 5}
		listaFloat    = []float64{1.0, 2, 0, 3.0, 4.0, 5.0}
		listaAny      = []any{1, 2, "3", 4, 5}
	)

	fmt.Println("---------- SEM GENERICS ----------")
	fmt.Println(SomarInteiros(listaInteiros))
	fmt.Println(SomarFloat(listaFloat))
	fmt.Println(SomarAny(listaAny))

	fmt.Println("---------- COM GENERICS ----------")
	fmt.Println(SomaGenerics(listaInteiros))
	fmt.Println(SomaGenerics(listaFloat))

}
