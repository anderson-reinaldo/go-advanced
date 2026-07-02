package main

import (
	"fmt"
	"time"
)

func main() {
	//Goroutines

	//duas ou mias goroutines acessando o mesmo dado ao mesmo tempo
	// pelo menos um desses acessos tem que ser uma escrita.

	m := make(map[int]int)

	for i := 0; i < 1000; i++ {
		m[i] = i
	}

	m2 := make(map[int]int)

	go func() {
		for i := 0; i < 1000; i++ {
			m2[i] = i
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(m2[i])
		}
	}()

	time.Sleep(time.Second * 5)

}
