package main

import "fmt"

func step1(ch1 chan<- int, numeros []int) {
	for _, numero := range numeros {
		ch1 <- numero
	}
	close(ch1)
}

func step2(ch1 <-chan int, ch2 chan<- int) {
	for numero := range ch1 {
		ch2 <- numero * 2
	}
	close(ch2)
}

func step3(ch2 <-chan int) {
	for result := range ch2 {
		fmt.Printf("Resultado recebido: %d\n", result)
	}
}

func main() {
	var (
		numeros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		ch1     = make(chan int)
		ch2     = make(chan int)
	)

	go step1(ch1, numeros)
	go step2(ch1, ch2)
	step3(ch2)

}
