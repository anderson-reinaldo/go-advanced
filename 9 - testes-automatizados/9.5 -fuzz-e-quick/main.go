package main

import (
	"fmt"

	"github.com/anderson-reinaldo/go-advanced/testes-automatizados/tests/fuzzy/utils"
)

func main() {
	invertida := utils.InverterString("Minha String")
	fmt.Println(invertida)

	slice := []int{1, 2, 3, 4, 5, 7, 10}
	fmt.Println(utils.VerificarOrdenacao(slice))
}
