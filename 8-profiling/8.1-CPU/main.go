package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Criação do arquivo de perfil (profile) de CPU
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Erro ao criar arquivo de perfil:", err)
		return
	}
	defer f.Close()

	//inicio do profiling de CPU
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("Erro ao iniciar perfil de CPU:", err)
		return
	}
	defer pprof.StopCPUProfile()

	//execulta uma carga de trabalho complexa
	trabalhoComplexo()

	fmt.Println("perfil de CPU salvo em 'cpu.prof'")
}

func trabalhoComplexo() {
	resultados := make(chan int, 5)

	//Execulta operações concorrentes
	for i := 0; i < 5; i++ {
		go func() {
			resultados <- somaPesada()
		}()
	}

	//Aguarda os resultados
	for i := 0; i < 5; i++ {
		resultado := <-resultados
		fmt.Printf("Resultado %d: %d\n", i+1, resultado)
	}

	//Realiza operações sequenciais
	for i := 0; i < 3; i++ {
		calculoPesado()
		fmt.Printf("Calculo pesado realizado. Interação: %d\n", i)
	}
}

func somaPesada() int {
	total := 0
	for i := 0; i < 1_000_000; i++ {
		total += rand.IntN(100)
	}

	return total
}

func calculoPesado() {
	time.Sleep(time.Millisecond * 500) //Simula I/O
	for i := 0; i < 500_000; i++ {
		_ = i * i
	}
}
