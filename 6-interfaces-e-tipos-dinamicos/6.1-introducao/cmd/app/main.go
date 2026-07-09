package main

import (
	"github.com/anderson-reinaldo/go-advanced/interfaces-e-tipos-dinamicos/introducao/pkg/conversor"
	"github.com/anderson-reinaldo/go-advanced/interfaces-e-tipos-dinamicos/introducao/pkg/conversor/json"
	"github.com/anderson-reinaldo/go-advanced/interfaces-e-tipos-dinamicos/introducao/pkg/conversor/xml"
)

func main() {
	gerenciador := conversor.NovoGerenciador()
	gerenciador.RegistrarConversor(json.ConversorJSON{})
	gerenciador.RegistrarConversor(xml.ConversorXML{})

	dadosJSON := `{
		"pessoa": {
		"nome": "Reinaldo",
		"idade": 25
		}
	}
	`

	dadosXML := `
		<pessoa>
			<nome>Reinaldo</nome>
			<idade>25</idade>
		</pessoa>
	`

	gerenciador.Processar("JSON", dadosJSON)
	gerenciador.Processar("XML", dadosXML)

}
