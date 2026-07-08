package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func testeMutex(interations int) (int64, time.Duration) {

	var (
		counter int64
		mu      sync.Mutex
		wg      sync.WaitGroup
		start   = time.Now()
	)

	for i := 0; i < interations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	return counter, time.Since(start)
}

func testeAtomic(interations int) (int64, time.Duration) {
	var (
		counter int64
		wg      sync.WaitGroup
		start   = time.Now()
	)

	for i := 0; i < interations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	return counter, time.Since(start)
}

func percent(t1, t2 time.Duration) float64 {
	//Assume t2 < t1 (porque o pool sempre é mais rapido)
	return (float64(t1-t2) / float64(t1)) * 100.0
}

func main() {

	interations := 1_000_000

	fmt.Printf("Execultando comparação com %d interações...\n", interations)

	c1, t1 := testeMutex(interations)

	fmt.Printf("[MUTEX] Contador Final %d, Tempo: %s\n", c1, t1)

	c2, t2 := testeAtomic(interations)

	fmt.Printf("[ATOMIC] Contador Final %d, Tempo: %s\n", c2, t2)

	fmt.Printf("[RESULTADO] %.2f%% mais rapido \n", percent(t1, t2))
}
