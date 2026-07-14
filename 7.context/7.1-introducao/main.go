package main

import (
	"context"
	"fmt"
)

type (
	myString1 string
	myString2 string
)

func main() {
	//O context no GO é uma ferramenta para passar informações compartilhadas e controlar o clico de vida de operações do progama.
	ctx := context.Background()
	ctx = context.WithValue(ctx, myString1("chave"), "valor")
	ctx = context.WithValue(ctx, myString2("chave"), "valor2")

	leitura(ctx)
}

func leitura(ctx context.Context) {
	valorArmazenado := ctx.Value(myString2("chave"))
	if valorArmazenado != nil {
		fmt.Println("Valor encontrado: ", valorArmazenado)
	} else {
		fmt.Println("Valor não foi encontrado.")
	}
}
