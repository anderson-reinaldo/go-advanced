package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"runtime/pprof"
	"time"

	"github.com/anderson-reinaldo/go-advanced/profiling/heap/cache"
)

func main() {

	//Simula vazamento de memoria
	go simularVazamento(cache.New())

	//Aguarda para capturar o perfil
	time.Sleep(3 * time.Second)

	// Criação do arquivo de perfil (profile) de Heap
	f, err := os.Create("heap.prof")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de perfil:", err)
		return
	}
	defer f.Close()

	//inicio do profiling de CPU
	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Println("Erro ao iniciar perfil de Heap:", err)
		return
	}
	defer pprof.StopCPUProfile()

	//execulta uma carga de trabalho complexa

	fmt.Println("perfil de Heap salvo em 'heap.prof'")
}

func simularVazamento(c *cache.Cache) {
	for i := 0; i < 1000; i++ {
		//Adiciona itens ao cache sem removê-los
		chave := fmt.Sprintf("chave-%d", rand.IntN(1_000_000))
		tamanho := rand.IntN(1000000) + 10000 //aloca entre 10KB e 100KB
		c.Inserir(chave, tamanho)

	}

	/*
		if len(c.Dados) > 10 {
			c.Limpar()
		}
	*/

	time.Sleep(10 * time.Millisecond) //Simula trabalho

}
