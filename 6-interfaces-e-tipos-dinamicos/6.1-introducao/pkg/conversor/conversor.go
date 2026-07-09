package conversor

import "fmt"

type Conversor interface {
	Converter(string) (map[string]any, error)
	Formato() string
}

type Gerenciador struct {
	conversores map[string]Conversor
}

func NovoGerenciador() *Gerenciador {
	return &Gerenciador{
		conversores: make(map[string]Conversor),
	}
}

func (g *Gerenciador) RegistrarConversor(c Conversor) {
	g.conversores[c.Formato()] = c
}

func (g *Gerenciador) Processar(formato, dados string) {
	conversor, ok := g.conversores[formato]
	if !ok {
		fmt.Printf("Conversor do formato %s não encontrado.\n", formato)
		return
	}

	resultado, err := conversor.Converter(dados)
	if err != nil {
		fmt.Printf("Erro ao converter formato %s\n", formato)
		return
	}

	fmt.Printf("Dados processados %v\n", resultado)

}
