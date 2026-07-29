package main

import (
	"reflect"
	"testing"
)

func TestContarPalavras(t *testing.T) {
	input := "olá mundo olá go go go"

	outputEsperado := map[string]int{"go": 3, "mundo": 1, "olá": 2} //want
	output := ContarPalavras(input)

	if !reflect.DeepEqual(output, outputEsperado) {
		t.Errorf("ContarPalavras(%v) = %v; esperado %v", input, output, outputEsperado)
	}
}
