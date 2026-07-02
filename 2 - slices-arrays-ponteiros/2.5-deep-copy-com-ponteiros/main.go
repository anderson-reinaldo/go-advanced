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

	pessoa2.Nome = toPointer("teste")

	fmt.Println("Nome da pessoa 1: ", *pessoa1.Nome)
	fmt.Println("Nome da pessoa 2: ", *pessoa2.Nome)

	fmt.Println("---------")

	pessoa3 := Pessoa{
		Nome:  toPointer("Anderson"),
		Idade: 25,
	}
	pessoa4 := deepCopy(pessoa3)
	pessoa4.Idade = 28

	fmt.Println("Idade da pessoa 3: ", pessoa4.Idade)

	fmt.Println("---------")
	listaDePessoas1 := []Pessoa{pessoa1, pessoa2, pessoa3, pessoa4}
	listaDePessoas2 := deepCopyList(listaDePessoas1)

	fmt.Println("Lista de pessoas 1: ", listaDePessoas1)
	fmt.Println("Lista de pessoas 2: ", listaDePessoas2)
}

func toPointer(s string) *string {
	return &s
}

func deepCopy(p Pessoa) Pessoa {
	var destino Pessoa

	destino.Idade = p.Idade
	destino.Nome = toPointer(*p.Nome)

	return destino
}

func deepCopyList(origem []Pessoa) []Pessoa {
	var destino = make([]Pessoa, len(origem))

	for i, pessoa := range origem {
		destino[i] = deepCopy(pessoa)
	}

	return destino
}
