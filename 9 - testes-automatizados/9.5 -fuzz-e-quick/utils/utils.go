package utils

import "unicode/utf8"

func InverterString(s string) string {
	if !utf8.ValidString(s) {
		return ""
	}

	runas := []rune(s) // fazemos isso para representar um caractere Unicode ao invés de uma string. Pois string é um []byte imutável
	for i, j := 0, len(runas)-1; i < j; i, j = i+1, j-1 {
		runas[i], runas[j] = runas[j], runas[i]
	}

	return string(runas)
}

func VerificarOrdenacao(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}

	return true
}
