package main

import (
	"fmt"
	"time"
)

func enviar(ch chan int) {
	fmt.Println("Enviando valor para o channel...")
	ch <- 100
	fmt.Println("Valor enviado com sucesso")
}

func receber(ch chan int) {
	time.Sleep(time.Second)
	valor := <-ch
	fmt.Println("Valor recebido: ", valor)
}

func main() {
	//Comunicação entre goroutines.
	ch := make(chan int) //Channel sem buffer

	//Channels sem buffer são canais bloqueantes o que quer dizer que se eu enviar um dado no canal eu so posso enviar outro quando o primeiro for lido
	// e se no caso eu tentar ler mais do que eu enviei caio em DEADLOCK

	for i := 0; i < 3; i++ {
		go enviar(ch)
	}

	time.Sleep(time.Second)

	for i := 0; i < 3; i++ {
		go receber(ch)
	}

	time.Sleep(time.Second * 5)

	//A linha abaixo causa um deadlock pois nunca vai ter um quarto envio de canais pois o for de envio so envia 3 vezes.
	//<-ch

}
