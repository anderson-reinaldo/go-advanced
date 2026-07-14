package main

import (
	"context"
	"fmt"
	"time"
)

type myString string

const (
	timestamp myString = "timestamp"
	id        myString = "id"
)

func escreverCTX(indentificador string, ctx context.Context) {
	fmt.Printf("[%-6s] timestamp=%v, id=%v\n", indentificador, ctx.Value(timestamp), ctx.Value(id))
}

func main() {

	ctxPai, cancelPai := context.WithCancel(
		context.WithValue(context.Background(), timestamp, time.Now().Format(time.RFC3339)),
	)

	ctxFilho1, cancelFilho1 := context.WithCancel(
		context.WithValue(ctxPai, id, "ctx-filho-1"),
	)

	ctxFilho2, _ := context.WithCancel(
		context.WithValue(ctxPai, id, "ctx-filho-2"),
	)

	fmt.Println("===== VALORES DENTRO DOS CONTEXTOS =====")

	escreverCTX("pai", ctxPai)
	escreverCTX("filho1", ctxFilho1)
	escreverCTX("filho2", ctxFilho2)

	go func() {
		<-ctxPai.Done()
		fmt.Println("Contexto pai cancelado: ", ctxPai.Err())
	}()

	go func() {
		<-ctxFilho1.Done()
		fmt.Println("Contexto filho 1 cancelado: ", ctxFilho1.Err())
	}()

	go func() {
		<-ctxFilho2.Done()
		fmt.Println("Contexto filho 2 cancelado: ", ctxFilho2.Err())
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Cancelando contexto filho 1...")
	cancelFilho1()

	time.Sleep(time.Second * 2)
	fmt.Println("Cancelando contexto pai...")
	cancelPai()

	time.Sleep(time.Second * 2)

}
