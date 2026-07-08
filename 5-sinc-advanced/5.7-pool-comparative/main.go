package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

func semPool(interations int) time.Duration {
	start := time.Now()

	for i := 0; i < interations; i++ {
		buffer := bytes.Buffer{}
		//Cria um novo buffer a cada interacao
		buffer.WriteString("Ola, mundo!")
		_ = buffer.String() //Simula o uso do buffer
	}

	return time.Since(start)
}

func comPool(interations int) time.Duration {
	start := time.Now()
	//Define o pool de buffers
	pool := sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	for i := 0; i < interations; i++ {
		//reaproveitando buffers
		buffer := pool.Get().(*bytes.Buffer)
		buffer.WriteString("Ola, mundo!")
		_ = buffer.String() //Simula o uso do buffer

		buffer.Reset()
		pool.Put(buffer)
	}

	return time.Since(start)
}

func percent(t1, t2 time.Duration) float64 {
	//Assume t2 < t1 (porque o pool sempre é mais rapido)
	return (float64(t1-t2) / float64(t1)) * 100.0
}

func main() {

	interations := 1_000_000

	fmt.Printf("Execultando comparação com %d interações...\n", interations)

	t1 := semPool(interations)

	fmt.Printf("[SEM POOL] Tempo total: %s\n", t1)

	t2 := comPool(interations)

	fmt.Printf("[COM POOL] Tempo total: %s\n", t2)

	fmt.Printf("[RESULTADO] %.2f%% mais rapido \n", percent(t1, t2))

}

/*
Execultando comparação com 1000000000 interações...
[SEM POOL] Tempo total: 1m28.915184194s
[COM POOL] Tempo total: 21.325983825s
[RESULTADO] 76.02% mais rapido
*/
