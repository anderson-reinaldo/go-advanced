package main

import (
	"fmt"
)

func runBroadCast(channels []chan string, message string) {
	for _, channel := range channels {
		channel <- message
	}
}

func main() {
	messages := []string{"Olá", "Mundo", "!"}
	var channels = []chan string{
		make(chan string),
		make(chan string),
		make(chan string),
	}

	for _, message := range messages {
		go runBroadCast(channels, message)
	}

	fmt.Println("Processo encerrado.")
}
