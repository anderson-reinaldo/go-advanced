package main

import (
	"fmt"
	"time"
)

func enviarMensagem(ch chan<- string, message string, sleep time.Duration) {
	for {
		time.Sleep(sleep)
		ch <- message
	}
}

func main() {
	var (
		ch1 = make(chan string)
		ch2 = make(chan string)
	)

	go enviarMensagem(ch1, "Enviando mensagem para o canal 1", time.Second*6)
	go enviarMensagem(ch2, "Enviando mensagem para o canal 2", time.Second*7)

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(time.Second * 5):
			fmt.Println("Nenhuma mensagem recebida")
		}

	}
}
