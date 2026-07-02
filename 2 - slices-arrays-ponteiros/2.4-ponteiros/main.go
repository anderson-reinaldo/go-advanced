package main

import "fmt"

type Pessoa struct {
	Nome     string
	Idade    int
	telefone *string
}

func main() {
	var a int = 10
	var b *int = &a

	fmt.Println("Valor de a: ", a)
	fmt.Println("Valor de b: ", b)
	fmt.Println("Valor de b desreferenciado: ", *b)

	p := NovaPessoa("Reinaldo", 25)
	p.AtualizarIdade(26)
	p.AtualizarTelefone("4002-8922")

	fmt.Println("NOVA PESSOA: ", p.Nome, p.Idade, p.Telefone())

}

func NovaPessoa(nome string, idade int) Pessoa {
	return Pessoa{
		Nome:  nome,
		Idade: idade,
	}
}

func (p *Pessoa) AtualizarIdade(idade int) {
	p.Idade = idade
}

func (p *Pessoa) AtualizarTelefone(telefone string) {
	p.telefone = &telefone
}

func (p Pessoa) Telefone() string {
	if p.telefone == nil {
		return ""
	}

	return *p.telefone
}
