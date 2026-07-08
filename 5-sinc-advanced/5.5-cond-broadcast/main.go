package main

import (
	"fmt"
	"sync"
	"time"
)

type Trava struct {
	liberada bool
	cond     *sync.Cond
	mu       sync.Mutex
}

func NovaTrava() *Trava {
	t := &Trava{}
	t.cond = sync.NewCond(&t.mu)
	return t
}

func worker(id int, t *Trava, wg *sync.WaitGroup) {
	defer wg.Done()

	t.mu.Lock()

	for !t.liberada {
		logf(fmt.Sprintf("Worker %d aguardando a trava liberar...", id))
		t.cond.Wait()
	}

	t.mu.Unlock()

	logf(fmt.Sprintf("Worker %d execultando...", id))
	time.Sleep(time.Millisecond * 100)
}

var start = time.Now()

func logf(msg string) {
	fmt.Printf("[%4dms] %s\n", time.Since(start).Milliseconds(), msg)
}

func main() {
	{
		logf("========== BROADCAST ==========")

		var (
			qtdeWorkers = 4
			trava       = NovaTrava()
			wg          sync.WaitGroup
		)

		wg.Add(qtdeWorkers)
		for i := 0; i < qtdeWorkers; i++ {
			go worker(i, trava, &wg)
		}

		time.Sleep(time.Millisecond * 500)
		trava.mu.Lock()
		trava.liberada = true
		logf("Trava Liberada!")
		trava.cond.Broadcast()
		trava.mu.Unlock()

		wg.Wait()
	}

	{
		logf("========== SIGNAL ==========")
		var (
			qtdeWorkers = 4
			trava       = NovaTrava()
			wg          sync.WaitGroup
		)

		wg.Add(qtdeWorkers)
		for i := 0; i < qtdeWorkers; i++ {
			go worker(i, trava, &wg)
		}

		time.Sleep(time.Millisecond * 300)
		trava.mu.Lock()
		trava.liberada = true
		logf("Chamando Signal() outra vez!")
		trava.cond.Signal()
		trava.mu.Unlock()

		time.Sleep(time.Millisecond * 300)
		trava.mu.Lock()
		trava.liberada = true
		logf("Chamando Signal() outra vez!")
		trava.cond.Signal()
		trava.mu.Unlock()

		time.Sleep(time.Millisecond * 300)
		trava.mu.Lock()
		trava.liberada = true
		logf("Chamando Signal() outra vez!")
		trava.cond.Signal()
		trava.mu.Unlock()

		time.Sleep(time.Millisecond * 300)
		trava.mu.Lock()
		trava.liberada = true
		logf("Chamando Signal() outra vez!")
		trava.cond.Signal()
		trava.mu.Unlock()

		wg.Wait()
	}
}
