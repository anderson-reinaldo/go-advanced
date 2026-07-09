package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type (
	Pessoa struct {
		Nome  string
		Idade int
	}

	Casa struct {
		Endereco string
		Cor      string
		Numero   int64
	}

	Buscavel interface {
		constraints.Integer | string | Pessoa | Casa
	}
)

//Crie uma funcao que seja capaz de buscar um item em uma lista de Pessoas(nome | idade), numeros e strings e retorne a posição desse item na lista

// [10,20,30] -> 20 -> 1

// Suporte qualquer tipo de inteiro, strigs, Pessoa

func buscar[T Buscavel](data []T, f func(T) bool) int {
	for idx, value := range data {
		if f(value) {
			return idx
		}
	}

	return -1
}

func main() {
	var (
		listaInteiros = []int{10, 20, 30, 40}
		listaStrings  = []string{"A", "B", "C", "D"}
		listaPessoas  = []Pessoa{
			{
				Nome:  "Anderson",
				Idade: 25,
			},
			{
				Nome:  "Reinaldo",
				Idade: 26,
			},
		}

		listaCasas = []Casa{
			{
				Endereco: "Rua chico matias",
				Cor:      "Branca",
				Numero:   56,
			},
			{
				Endereco: "Rua J",
				Cor:      "Verde",
			},
		}
	)

	fmt.Printf("O numero 30 está na posição: %d\n", buscar(listaInteiros, func(n int) bool {
		return n == 30
	}))
	fmt.Printf("O texto B está na posição: %d\n", buscar(listaStrings, func(s string) bool {
		return s == "B"
	}))
	fmt.Printf("A pessoa com nome Anderson está na posição: %d\n", buscar(listaPessoas, func(p Pessoa) bool {
		return p.Nome == "Anderson"
	}))
	fmt.Printf("A pessoa com idade 26 está na posição: %d\n", buscar(listaPessoas, func(p Pessoa) bool {
		return p.Idade == 26
	}))

	fmt.Printf("A Casa com numero 10 está na posição: %d\n", buscar(listaCasas, func(p Casa) bool {
		return p.Numero == 10
	}))

	fmt.Printf("A Casa com endereco rua chico matias está na posição: %d\n", buscar(listaCasas, func(p Casa) bool {
		return p.Endereco == "Rua chico matias"
	}))

}
