package main

import (
	"fmt"

	"github.com/anderson-reinaldo/go-advanced/testes-automatizados/tests/benchmarking/math"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(math.SomaIterativa(numeros))
	fmt.Println(math.SomaRecursiva(numeros))
}
