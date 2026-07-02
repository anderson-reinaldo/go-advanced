package main

import "fmt"

func main() {
	// Estrutura de chave-valor

	var meuMap = map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
	}

	/*
	   +----------+----------------------+
	   | Bucket 0 | [("a", 1), ("e", 5)] |
	   +----------+----------------------+
	   | Bucket 1 | [("b", 2)]           |
	   +----------+----------------------+
	   | Bucket 2 | [("c", 3), ("f", 6)] |
	   +----------+----------------------+
	   | Bucket 3 | [("d", 4)]           |
	   +----------+----------------------+
	*/

	for key, value := range meuMap {
		fmt.Println(key, value)
	}

	var meuMap2 = map[string]any{}
	meuMap2["chave"] = 1

	var meuValor, ok = meuMap2["chave"].(string)
	if ok {
		fmt.Println(meuValor)
	} else {
		fmt.Println("Valor não encontrado.")
	}

}
