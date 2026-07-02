package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  = make(map[int]int)
		mu sync.Mutex
		wg sync.WaitGroup
	)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			//trava execulcoes que podem causar concorrencia e a proxima goroutine espera o Unlock() ser execultado
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Fim da execulção", m)

}
