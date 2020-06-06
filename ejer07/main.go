package main

import "fmt"

var Calculo = func(num1 int, num2 int) int {
	return num1 + num2

}

func main() {
	result := Calculo(5, 10)
	fmt.Println(result)
	//restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	result = Calculo(15, 5)
	fmt.Println(result)

	Operaciones()
	//Closures
	tableDel := 2
	MiTable := Tabla(tableDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTable())
	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
