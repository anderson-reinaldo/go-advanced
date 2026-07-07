package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	dados map[string]any
	mu    sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		dados: make(map[string]any),
	}
}

func (c *Cache) Set(chave string, valor any) {
	logf("Set: Tentativa de Escrita (Lock)")
	c.mu.Lock()
	time.Sleep(time.Millisecond * 300)
	logf("Set: Lock Adquirido")
	c.dados[chave] = valor
	defer c.mu.Unlock()

	logf("Set: Unlock (escrita finalizada)")
}

func (c *Cache) Get(chave string) (any, bool) {
	logf("Get: Tentativa de leitura (RLock)")
	c.mu.RLock()
	logf("Get: RLock Adquirido")
	time.Sleep(time.Millisecond * 100)
	defer c.mu.RUnlock()
	logf("Get: RUnlock (Leitura finalizada)")

	valor, ok := c.dados[chave]
	if ok {
		return valor, true
	} else {
		return nil, false
	}
}

var start = time.Now()

func logf(msg string) {
	fmt.Printf("[%4dms] %s\n", time.Since(start).Milliseconds(), msg)
}

func main() {
	var (
		cache = NewCache()
		wg    sync.WaitGroup
	)

	cache.dados["k"] = "v"
	wg.Add(4)

	go func() {
		defer wg.Done()
		logf("Goroutine 1: chama Get()")
		_, _ = cache.Get("k")
		logf("Goroutine 1: recebeu retorno do Get()")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 10)
		logf("Goroutine 2: chama Get()")
		_, _ = cache.Get("k")
		logf("Goroutine 2: recebeu retorno do Get()")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 20)
		logf("Goroutine 3: chama Set() - (Quer Lock))")
		cache.Set("k", "v2")
		logf("Goroutine 3: recebeu retorno do Set()")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 300)
		logf("Goroutine 4: chama Get()")
		_, _ = cache.Get("k")
		logf("Goroutine 4: recebeu retorno do Get()")
	}()

	wg.Wait()

	logf("FIM")

}
