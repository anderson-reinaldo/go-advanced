package main

import (
	"context"
	"fmt"
	"time"

	"github.com/anderson-reinaldo/go-advanced/context/contexto-customizado/custom"
)

func main() {
	//Contexto base com timeout
	ctxPai, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	//Cria um contexto customizado que ignora o timeout
	customCtx := custom.NewContext(ctxPai, map[string]any{
		"chave": "valor",
	})

	//Imagine uma goroutine que verifica o comportamento do CustomContext
	go func() {
		select {
		case <-customCtx.Done(): //Não será execultado, pois o Done do pai não influencia
			fmt.Println("Contexto customizado foi cancelado...", customCtx.Err())
		case <-time.After(5 * time.Second): //Simula trabalho longo
			fmt.Println("5 Segundos se passaram...")
		}
	}()

	time.Sleep(6 * time.Second)
	fmt.Println("Finalizado")
}
