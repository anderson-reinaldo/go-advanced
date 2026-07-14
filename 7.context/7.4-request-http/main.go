package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func fazerRequest(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	//time.Sleep(time.Second * 6)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	url := "https://jsonplaceholder.typicode.com/posts/1"
	dados, err := fazerRequest(ctx, url)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Timeout na execução")
		} else if ctx.Err() == context.Canceled {
			fmt.Println("Contexto cancelado manualmente!")
		} else {
			fmt.Println("Erro:", err.Error())
		}
		return
	} else {
		fmt.Println("Resultado: ", dados)
	}
}
