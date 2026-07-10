package main

import (
	"fmt"
)

type (
	Cache[K comparable, V any] struct {
		dados map[K]V
	}
)

func NovoCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		dados: make(map[K]V),
	}
}

func (c *Cache[K, V]) Set(chave K, valor V) {
	c.dados[chave] = valor
}

func (c *Cache[K, V]) Get(chave K) (V, bool) {
	valor, ok := c.dados[chave]
	return valor, ok
}

func main() {
	cache := NovoCache[string, int]()
	cache.Set("key-1", 1)

	v, ok := cache.Get("key-1")
	if ok {
		fmt.Printf("O valor de key-1 é: %d\n", v)
	}

	cache2 := NovoCache[int, string]()
	cache2.Set(0, "valor 1")

	v2, ok := cache2.Get(0)
	if ok {
		fmt.Printf("O valor de 0 é: %s\n", v2)
	}

}
