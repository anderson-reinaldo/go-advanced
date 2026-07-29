package math

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func gerarSlice(tamanho int) []int {
	slice := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		slice[i] = rand.IntN(100)
	}

	return slice
}

func BenchmarkSomaIterativa(b *testing.B) {
	// gera os dados de teste UMA vez só
	numeros := gerarSlice(10)

	// ResetTimer zera o contador de tempo do benchmark.
	// Tudo que veio antes (gerarSlice) não entra na medição.
	b.ResetTimer()

	// O framework de testes define b.N automaticamente.
	// Ele vai aumentando b.N até ter confiança no tempo médio por operação.
	for i := 0; i < b.N; i++ {
		_ = SomaIterativa(numeros) // chamamos a função que queremos medir N vezes
	}
}

func BenchmarkSomaRecursiva(b *testing.B) {
	// gera os dados de teste UMA vez só
	numeros := gerarSlice(10)

	// ResetTimer zera o contador de tempo do benchmark.
	// Tudo que veio antes (gerarSlice) não entra na medição.
	b.ResetTimer()

	// O framework de testes define b.N automaticamente.
	// Ele vai aumentando b.N até ter confiança no tempo médio por operação.
	for i := 0; i < b.N; i++ {
		_ = SomaRecursiva(numeros) // chamamos a função que queremos medir N vezes
	}
}

// Testa diferentes tamanhos de slice
func BenchmarkSomaIterativaTamanhosDiferentes(b *testing.B) {
	tamanhos := []int{100, 1000, 10000, 100000, 1_000_000}

	for _, tamanho := range tamanhos {
		nome := fmt.Sprintf("Tamanho - %d", tamanho)

		b.Run(nome, func(b *testing.B) {
			numeros := gerarSlice(tamanho)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = SomaIterativa(numeros)
			}
		})
	}
}

// Testa diferentes tamanhos de slice
func BenchmarkSomaRecursivaTamanhosDiferentes(b *testing.B) {
	tamanhos := []int{100, 1000, 10000, 100000, 1_000_000}

	for _, tamanho := range tamanhos {
		nome := fmt.Sprintf("Tamanho - %d", tamanho)

		b.Run(nome, func(b *testing.B) {
			numeros := gerarSlice(tamanho)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = SomaRecursiva(numeros)
			}
		})
	}
}
