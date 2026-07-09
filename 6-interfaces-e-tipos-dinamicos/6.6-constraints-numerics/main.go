package main

import "fmt"

type Myint int

type Number interface {
	~int | float64
}

func Soma[T Number](a, b T) T {
	return a + b
}

func main() {
	var (
		x1 = int(100)
		x2 = int(200)

		y1 = float64(100.5)
		y2 = float64(200.5)

		z1 Myint = 100
		z2 Myint = 200
	)

	fmt.Println("Somando x1 + x2: ", Soma(x1, x2))
	fmt.Println("Somando y1 + y2: ", Soma(y1, y2))

	/*
		Esse caso abaixo nao pode pos nao pode somar numeros com tipos diferentes e o GO por baixo dos panos
		infere todos os tipos pelo valor do primeiro paramentro passado que o tipo generico aceita.
		--> fmt.Println("Somando x1 + y1: ", Soma(x1, y1))
	*/

	//	fmt.Println("Somando x1 + x2 convertendo para int64: ", Soma(int64(x1), int64(x2)))
	fmt.Println("Somando z1 + z2: ", Soma(z1, z2))

}
