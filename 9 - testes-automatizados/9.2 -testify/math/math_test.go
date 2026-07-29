package math

import (
	"testing"

	"github.com/anderson-reinaldo/go-advanced/testes-automatizados/tests/testify-lib/erros"
	"github.com/stretchr/testify/assert"
)

func TestMedia(t *testing.T) {
	testes := []struct {
		nome           string
		input          []float64
		outputEsperado float64
		erroEsperado   error
	}{
		{
			nome:           "lista normal",
			input:          []float64{10, 20, 30},
			outputEsperado: 20,
			erroEsperado:   nil,
		},
		{
			nome:           "lista vazia",
			input:          []float64{},
			outputEsperado: 0,
			erroEsperado:   erros.ErroListaVazia,
		},
		{
			nome:           "todos iguais",
			input:          []float64{5, 5, 5, 5},
			outputEsperado: 5,
			erroEsperado:   nil,
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			output, err := Media(tt.input)

			if tt.erroEsperado != nil {
				assert.ErrorIs(
					t,
					err,
					tt.erroEsperado,
					"era esperado erro %v para entrada %v",
					tt.erroEsperado,
					tt.input,
				)

			} else {
				assert.NoError(
					t,
					err,
					"não era esperado erro para entrada %v",
					tt.input,
				)
			}

			assert.Equal(
				t,
				tt.outputEsperado,
				output,
				"Media(%v) retornou %v; esperado %v",
				tt.input,
				output,
				tt.outputEsperado,
			)
		})
	}

}

func TestMDC(t *testing.T) {
	testes := []struct {
		nome           string
		n1             int
		n2             int
		outputEsperado int
		erroEsperado   error
	}{
		{
			nome:           "MDC comum (48, 18)",
			n1:             48,
			n2:             18,
			outputEsperado: 6,
			erroEsperado:   nil,
		},
		{
			nome:           "números primos entre si (101, 103)",
			n1:             101,
			n2:             103,
			outputEsperado: 1,
			erroEsperado:   nil,
		},
		{
			nome:           "número negativo deve retornar erro",
			n1:             -10,
			n2:             5,
			outputEsperado: 0,
			erroEsperado:   erros.ErroNumeroNegativo,
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			resultado, err := MDC(tt.n1, tt.n2)

			if tt.erroEsperado != nil {
				assert.ErrorIs(
					t,
					err,
					tt.erroEsperado,
					"era esperado erro %v para MDC(%d, %d)",
					tt.erroEsperado,
					tt.n1,
					tt.n2,
				)
			} else {
				assert.NoError(
					t,
					err,
					"não era esperado erro para MDC(%d, %d)",
					tt.n1,
					tt.n2,
				)
			}

			assert.Equal(
				t,
				tt.outputEsperado,
				resultado,
				"MDC(%d, %d) = %d; esperado %d",
				tt.n1,
				tt.n2,
				resultado,
				tt.outputEsperado,
			)
		})
	}
}

func TestFibonacci(t *testing.T) {
	testes := []struct {
		nome           string
		numero         int
		outputEsperado []int
		erroEsperado   error
	}{
		{
			nome:           "sequência com 5 elementos",
			numero:         5,
			outputEsperado: []int{1, 1, 2, 3, 5},
			erroEsperado:   nil,
		},
		{
			nome:           "sequência vazia (0 elementos)",
			numero:         0,
			outputEsperado: []int{},
			erroEsperado:   nil,
		},
		{
			nome:           "sequência com 10 elementos",
			numero:         10,
			outputEsperado: []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
			erroEsperado:   nil,
		},
		{
			nome:           "número negativo deve retornar erro",
			numero:         -1,
			outputEsperado: nil,
			erroEsperado:   erros.ErroNumeroNegativo,
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			result, err := Fibonacci(tt.numero)

			if tt.erroEsperado != nil {
				assert.ErrorIs(
					t,
					err,
					tt.erroEsperado,
					"era esperado erro %v para Fibonacci(%d)",
					tt.erroEsperado,
					tt.numero,
				)
			} else {
				assert.NoError(
					t,
					err,
					"não era esperado erro para Fibonacci(%d)",
					tt.numero,
				)
			}

			// testify faz deep equal em slices, então dá pra usar Equal direto
			assert.Equal(
				t,
				tt.outputEsperado,
				result,
				"Fibonacci(%d) = %v; esperado %v",
				tt.numero,
				result,
				tt.outputEsperado,
			)
		})
	}
}
