package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

// instalação necessaria: go get golang.org/x/sync@latest

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker WG %d iniciado\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker WG %d finalizado\n", id)

}

func workerWithErrorGroup(id int) error {
	fmt.Printf("Worker EG %d iniciado\n", id)
	time.Sleep(time.Second)

	if id > 2 {
		return errors.New("Encontramos um valor maior que 2!")
	}

	fmt.Printf("Worker EG %d finalizado\n", id)

	return nil

}

func main() {
	fmt.Println("------WAIT GROUP-----")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("------FIM WAIT GROUP-----")

	fmt.Println("------ERROR GROUP-----")
	var eg errgroup.Group

	eg.SetLimit(3)

	for i := 1; i <= 5; i++ {
		eg.Go(func() error {
			return workerWithErrorGroup(i)
		})
	}

	err := eg.Wait()
	if err != nil {
		fmt.Printf("Encontramos um erro ao execultar as goroutines: %s", err)
	}

	fmt.Println("------FIM ERROR GROUP-----")

}
