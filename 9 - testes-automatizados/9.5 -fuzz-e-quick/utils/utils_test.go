package utils

import (
	"sort"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

func TestInverterString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		nome           string
		input          string
		outputEsperado string
	}{
		{
			nome:           "String vazia",
			input:          "",
			outputEsperado: "",
		},
		{
			nome:           "Caractere único",
			input:          "a",
			outputEsperado: "a",
		},
		{
			nome:           "Palíndromo",
			input:          "madam",
			outputEsperado: "madam",
		},
		{
			nome:           "String normal",
			input:          "hello",
			outputEsperado: "olleh",
		},
		{
			nome:           "String com espaços",
			input:          "hello world",
			outputEsperado: "dlrow olleh",
		},
		{
			nome:           "String com caracteres especiais",
			input:          "123!@#",
			outputEsperado: "#@!321",
		},
		{
			nome:           "String com caracteres unicode",
			input:          "こんにちは",
			outputEsperado: "はちにんこ",
		},
		{
			nome:           "String inválida",
			input:          "\xff",
			outputEsperado: "",
		},
	}

	for _, test := range tests {
		t.Run(test.nome, func(t *testing.T) {
			t.Parallel()
			result := InverterString(test.input)
			assert.Equal(t, test.outputEsperado, result)
		})
	}
}

func FuzzInverterString(f *testing.F) {
	// A ideia é testar uma regra geral que deve ser verdadeira pra qualquer string válida.

	// Casos iniciais para guiar o fuzzing. Esses valores SEMPRE serão testados, e servem como "ponto de partida"
	f.Add("hello")
	f.Add("12345")
	f.Add("a😊b")

	// Fuzz registra uma função que será chamada muitas vezes com strings geradas automaticamente.
	f.Fuzz(func(t *testing.T, input string) {

		t.Logf("Testando input: %q", input)

		// Chama a função a ser testada com a string gerada
		inputInvertido := InverterString(input)
		if inputInvertido == "" {
			t.Skip() // sinaliza "não é um caso interessante" (no nosso caso, não é uma string válida)
		}

		duplaInversao := InverterString(inputInvertido)
		assert.Equal(t, input, duplaInversao)

		// Se isso falhar, o fuzzer vai parar, imprimir o input que quebrou
		// e gravar esse caso como parte do corpus (pra futuras execuções).
	})
}

func TestVerificarOrdenacaoQuick(t *testing.T) {
	// aqui a ideia é tentar "provar" que essa função SEMPRE retorna true para entradas ordenadas aleatórias.

	err := quick.Check(func(slice []int) bool {
		// Chama a função a ser testada com um slice aleatório, depois de orderná-lo.
		sort.Ints(slice)
		return VerificarOrdenacao(slice)
	}, nil)

	assert.NoError(t, err, "slice ordenado não passou no teste")
}
