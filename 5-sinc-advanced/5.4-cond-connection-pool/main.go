package main

import (
	"fmt"
	"sync"
	"time"
)

type PoolDeConexoes struct {
	conexoes []int
	mu       sync.Mutex
	cond     *sync.Cond
}

func NovoPoolDeConexoes(tamanho int) *PoolDeConexoes {
	pool := &PoolDeConexoes{
		conexoes: make([]int, tamanho),
	}

	for i := 0; i < tamanho; i++ {
		pool.conexoes[i] = i + 1
	}

	pool.cond = sync.NewCond(&pool.mu)

	return pool
}

func (p *PoolDeConexoes) ObterConexao() int {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.conexoes) == 0 {
		fmt.Println("Nenhuma conexao disponivel. Aguardando...")
		p.cond.Wait()
	}

	conexao := p.conexoes[0]
	p.conexoes = p.conexoes[1:]
	fmt.Printf("Conexão %d adquirida.\n", conexao)

	return conexao

}

func (p *PoolDeConexoes) LiberarConexao(conexao int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.conexoes = append(p.conexoes, conexao)
	fmt.Printf("Conexão %d liberada.\n", conexao)

	p.cond.Signal()
}

func main() {
	poolDeConexoes := NovoPoolDeConexoes(2)

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			conexao := poolDeConexoes.ObterConexao()
			fmt.Printf("Goroutine %d usando a conexão %d.\n", i, conexao)
			time.Sleep(time.Second * 2)
			poolDeConexoes.LiberarConexao(conexao)

		}()
	}

	wg.Wait()
	fmt.Println("Fim")
}
