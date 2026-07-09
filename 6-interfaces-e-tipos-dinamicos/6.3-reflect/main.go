package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x any = 10

	t := reflect.TypeOf(x)
	fmt.Println("type de x: ", t)

	v := reflect.ValueOf(x)
	if v.CanInt() {
		xInt := v.Int()
		fmt.Println("Valor de x:", xInt)
	} else {
		fmt.Println("x não é um valor inteiro.")
	}

	fmt.Println("---------------------------")
	var y any = 10

	v2 := reflect.ValueOf(&y).Elem()
	v2.Set(reflect.ValueOf(10000))

	fmt.Println("Valor de Y: ", y)
	fmt.Println("Valor de V2: ", v2.Elem())

}
