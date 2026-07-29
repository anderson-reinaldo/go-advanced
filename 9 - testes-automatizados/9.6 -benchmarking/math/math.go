package math

func SomaIterativa(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}

	return soma
}

func SomaRecursiva(numeros []int) int {
	if len(numeros) == 0 {
		return 0
	}

	return numeros[0] + SomaRecursiva(numeros[1:])
}
