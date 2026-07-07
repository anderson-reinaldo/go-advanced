package main

import (
	"fmt"
	"sync"
	"time"
)

type Buffer struct {
	dados []int
	cond  *sync.Cond
	mu    sync.Mutex
}

func main() {
	buffer := Buffer{
		dados: make([]int, 0),
	}

	buffer.cond = sync.NewCond(&buffer.mu)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			buffer.mu.Lock()
			fmt.Printf("Produzindo item %d\n", i)
			buffer.dados = append(buffer.dados, i)
			buffer.cond.Signal()
			buffer.mu.Unlock()
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			buffer.mu.Lock()
			for len(buffer.dados) == 0 {
				buffer.cond.Wait()
			}

			item := buffer.dados[0]
			buffer.dados = buffer.dados[1:]
			fmt.Printf("Consumindo item %d\n", item)
			buffer.mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("Fim")

}
