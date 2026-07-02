package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	dados sync.Map
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(chave string, valor any) {
	c.dados.Store(chave, valor)
	fmt.Printf("Valor %x armazenado no cache com a chave %s\n", valor, chave)
}

func (c *Cache) Get(chave string) (any, bool) {
	return c.dados.Load(chave)
}

func main() {
	var (
		cache = NewCache()
		wg    sync.WaitGroup
	)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Set(fmt.Sprintf("key-%d", i), i)
		}()
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
			valor, ok := cache.Get(fmt.Sprintf("key-%d", i))
			if ok {
				fmt.Printf("Valor encontrado na key-%d: %v\n", i, valor)
			} else {
				fmt.Println("Não foi possivel ler o valor no cache!", i)
			}
		}()
	}

	wg.Wait()

}
