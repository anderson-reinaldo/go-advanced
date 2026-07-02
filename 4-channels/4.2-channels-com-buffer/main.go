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
	ch := make(chan int, 3) //Channel com buffer

	//Channels com buffer são canais bloqueantes porem so bloqueiam quando o numero de envios para o canal ultrapassa o tamanho do buffer (3)
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
