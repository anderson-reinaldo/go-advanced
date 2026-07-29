package usecase

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/anderson-reinaldo/go-advanced/testes-automatizados/tests/mini-project/repository"
)

func BuscarDados(client repository.HTTPClient, url string) (map[string]any, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status code inesperado")
	}

	var response map[string]any

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, errors.New("json invalido")
	}

	return response, nil
}

func CriarDados(client repository.HTTPClient, url string, data map[string]any) (map[string]any, error) {
	resp, err := client.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status code inesperado")
	}

	var response map[string]any

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, errors.New("json invalido")
	}

	return response, nil
}
