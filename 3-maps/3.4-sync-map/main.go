package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		m  sync.Map
		wg sync.WaitGroup
	)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			//trava execulcoes que podem causar concorrencia e a proxima goroutine espera o Unlock() ser execultado
			m.Store(fmt.Sprintf("key-%d", i), i)
		}()
	}

	wg.Wait()

	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

	valor, ok := m.Load("key-50")
	if ok {
		fmt.Println(valor)
	} else {
		fmt.Println("Valor não encontrado.")
	}

	fmt.Println("Fim da execulção")

}
