package main

import "fmt"

func main() {
	fmt.Println(Calculo(5, 50))
	fmt.Println(Calculo(1, 10))
	fmt.Println(Calculo(5))
	fmt.Println(Calculo(5, 50, 10, 2, 30))
}

func multiplicar(numero int) (z int) {
	z = numero * 2
	return
}

func test(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func Calculo(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}
	return total
}
