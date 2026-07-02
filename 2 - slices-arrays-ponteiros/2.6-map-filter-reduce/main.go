package main

import "fmt"

type mySlice []int

func main() {
	//map -> [1,2,3,4,5,6] --> [2,4,6,8,10,12]
	//filter -> [1,2,3,4,5,6] --> [2,4,6]
	//reduce -> [1,2,3,4,5,6] --> 21

	var lista mySlice = mySlice{1, 2, 3, 4, 5, 6}
	lista = lista.Filter(func(num int) bool {
		return num%2 == 0
	})

	lista = lista.Map(func(num int) int {
		return num * 2
	})

	result1 := lista.Reduce(func(i1, i2 int) int {
		return i1 + i2
	}, 0)

	fmt.Println("Resultado Sequencial: ", result1)

	var lista2 mySlice = mySlice{1, 2, 3, 4, 5, 6}

	result2 := lista2.Filter(func(num int) bool {
		return num%2 == 0
	}).Map(func(num int) int {
		return num * 2
	}).Reduce(func(i1, i2 int) int {
		return i1 + i2
	}, 0)

	fmt.Println("Resultado encadeado: ", result2)

}

func (ms mySlice) Filter(cond func(num int) bool) mySlice {
	var result mySlice

	for _, value := range ms {
		if cond(value) {
			result = append(result, value)
		}
	}

	return result
}

func (ms mySlice) Map(trans func(num int) int) mySlice {
	var result mySlice

	for _, value := range ms {
		result = append(result, trans(value))

	}

	return result
}

func (ms mySlice) Reduce(acc func(int, int) int, inicial int) int {
	var result int = inicial

	for _, value := range ms {
		result = acc(result, value)
	}

	return result
}
