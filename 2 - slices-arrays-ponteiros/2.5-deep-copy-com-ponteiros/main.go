package main

import "fmt"

type Pessoa struct {
	Nome  *string
	Idade int
}

func main() {
	nome := "reinaldo"
	pessoa1 := Pessoa{
		Nome:  &nome,
		Idade: 25,
	}

	pessoa2 := pessoa1

	fmt.Println("Nome da pessoa 1: ", *pessoa1.Nome)
	fmt.Println("Nome da pessoa 2: ", *pessoa2.Nome)

}
