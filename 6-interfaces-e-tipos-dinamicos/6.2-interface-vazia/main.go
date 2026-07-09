package main

import "fmt"

func main() {
	fmt.Println("String", 1, true)
	fmt.Println(1)
	fmt.Println(true)
	fmt.Println([]int{1, 2, 3, 4})

	// -----

	var (
		numero  int = 5
		numero2 any = 5
	)

	numero = numero + 1
	numero++
	// -----

	/*
		Abaixo vai gerar um erro pois é impossivel somar 1 com any
		numero2 = numero2 + 1
		numero2++
	*/

	//Fazendo a declaracao de tipo "casting"
	numero2Int, ok := numero2.(int)
	if ok {
		numero2Int++
	}

	// -----

	/*
		var (
			lista1 any = []int{1, 2, 3, 4, 5, 6}
			lista2 any = []int{1, 2, 3, 4, 5, 6}
		)

		if lista1 == lista2 {
			fmt.Println("Comparação impossivel.")
		}
	*/
	// -----

	var x any
	x = "string"
	x = 123
	x = true

	switch v := x.(type) {
	case int:
		fmt.Println("Inteiro", v)
	case string:
		fmt.Println("String", v)
	case float64:
		fmt.Println("Float64", v)
	case bool:
		fmt.Println("Booleano", v)
	}

}
