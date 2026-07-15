package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var mu sync.Mutex

func main() {
	// Ativa o profiling de block
	runtime.SetBlockProfileRate(1)

	//Simula contenção proposital de block
	for i := 0; i < 20; i++ {
		go simularBloqueio(i)
	}

	//Aguarda as goroutines bloquearem
	time.Sleep(3 * time.Second)

	//Gera o perfil de goroutines progamaticamente
	f, err := os.Create("block.prof")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de perfil:", err)
		return
	}
	defer f.Close()

	if err := pprof.Lookup("block").WriteTo(f, 0); err != nil {
		fmt.Println("Erro ao escrever perfil de block:", err)
		return
	}

	fmt.Println("Perfil de block salvo em 'block.prof'")
}

func simularBloqueio(id int) {
	for i := 0; i < 10; i++ {
		mu.Lock()
		fmt.Printf("Goroutine %d com lock...\n", id)
		time.Sleep(20 * time.Millisecond)
		mu.Unlock()
	}
}
