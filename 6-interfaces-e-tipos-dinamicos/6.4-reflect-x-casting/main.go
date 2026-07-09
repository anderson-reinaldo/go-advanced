package main

import (
	"fmt"
	"reflect"
	"time"
)

func castPuro(interacoes int, valor any) time.Duration {
	start := time.Now()

	for i := 0; i < interacoes; i++ {
		if _, ok := valor.(int); ok {
			_ = 2
		}
	}

	return time.Since(start)

}

func castComReflect(interacoes int, valor any) time.Duration {
	start := time.Now()

	for i := 0; i < interacoes; i++ {
		v := reflect.ValueOf(valor)
		if v.Kind() == reflect.Int {
			_ = int(v.Int())
		}
	}

	return time.Since(start)

}

func percent(t1, t2 time.Duration) float64 {
	//Assume t2 < t1 (porque o pool sempre é mais rapido)
	return (float64(t1-t2) / float64(t1)) * 100.0
}

func main() {
	var (
		interacoes         = 100_000_000
		numeroValido   any = 42
		numeroInvalido any = "não é int"
	)

	fmt.Printf("Execultando comparação com %d interacões...\n", interacoes)

	t1 := castComReflect(interacoes, numeroValido)
	fmt.Printf("[Cast com Reflect - Numero Valido] Tempo Total: %s\n", t1)

	t2 := castPuro(interacoes, numeroValido)
	fmt.Printf("[Cast Puro - Numero Valido] Tempo Total: %s\n", t2)

	fmt.Printf("[Resultado] %.2f%% mais rápido\n", percent(t1, t2))

	fmt.Println("-------------------------------------------")

	t3 := castComReflect(interacoes, numeroInvalido)
	fmt.Printf("[Cast com Reflect - Numero Invalido] Tempo Total: %s\n", t3)

	t4 := castPuro(interacoes, numeroInvalido)
	fmt.Printf("[Cast Puro - Numero Invalido] Tempo Total: %s\n", t4)

	fmt.Printf("[Resultado] %.2f%% mais rápido\n", percent(t1, t2))

}
