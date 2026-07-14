package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d encerrando...\n", id)
			return
		default:
			fmt.Printf("Worker %d trabalhando...\n", id)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	//Funciona melhor se você tiver uma data especifica para cancelar o contexto.
	//context.WithDeadline(ctx, time.Now())

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(time.Second * 4)
}
