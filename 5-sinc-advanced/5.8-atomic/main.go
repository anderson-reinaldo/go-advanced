package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		flag int64
		wg   sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Escrevendo na variavel...")
		time.Sleep(time.Second * 3)
		atomic.StoreInt64(&flag, 1)
		fmt.Println("Valor escrito com sucesso")
	}()

	go func() {
		defer wg.Done()

		for atomic.LoadInt64(&flag) == 0 {
			fmt.Println("O valor ainda esta zero para flag...")
			time.Sleep(time.Millisecond * 500)
		}

		fmt.Printf("Valor obtido para flag: %d\n", atomic.LoadInt64(&flag))
	}()

	wg.Wait()

	fmt.Println("Fim")
}
