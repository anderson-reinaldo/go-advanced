package main

import (
	"fmt"
	"sync"
)

// Pega todos os valores e joga em um canal
func produtor(numeros []int, entrada chan<- int) {
	for _, numero := range numeros {
		entrada <- numero
	}

	close(entrada)
}

// Pega o canal com todos os valores, faz o processamento necessesario e joga o valor dentro do canal de saida
func fanOut(id int, entrada <-chan int, saida chan<- int) {
	for numero := range entrada {
		fmt.Printf("Channel %d processando o numero %d\n", id, numero)
		saida <- numero * numero
	}
	close(saida)
}

// Pega todos os canais de saida, pega os valores e retorna apenas em um unico canal de saida
func fanIn(saidasReadOnly []<-chan int) <-chan int {
	var (
		saidaFinal = make(chan int)
		wg         sync.WaitGroup
	)

	wg.Add(len(saidasReadOnly))

	for _, saida := range saidasReadOnly {
		go func() {
			defer wg.Done()
			for numero := range saida {
				saidaFinal <- numero
			}
		}()
	}

	go func() {
		wg.Wait()
		close(saidaFinal)
	}()

	return saidaFinal
}

func main() {
	var (
		numeros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //Simulação de workload(carga de dados para tarefa)

		entrada = make(chan int)

		saidas = []chan int{
			make(chan int),
			make(chan int),
			make(chan int),
		}
	)

	go produtor(numeros, entrada)

	for i, saida := range saidas {
		go fanOut(i+1, entrada, saida)
	}

	saidasReadOnly := make([]<-chan int, len(saidas)) // cria uma lista de canais somente leitura
	for i, saida := range saidas {
		saidasReadOnly[i] = saida
	}

	saidaFinal := fanIn(saidasReadOnly)

	for resultado := range saidaFinal {
		fmt.Printf("Resultado recebido na saida final: %d\n", resultado)
	}
}
