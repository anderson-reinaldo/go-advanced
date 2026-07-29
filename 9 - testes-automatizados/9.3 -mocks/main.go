package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anderson-reinaldo/go-advanced/testes-automatizados/tests/mini-project/usecase"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	response, err := usecase.BuscarDados(http.DefaultClient, url)
	if err != nil {
		log.Fatal(err)
	}

	responseLegivel, _ := json.Marshal(response)

	fmt.Println(string(responseLegivel))
}
