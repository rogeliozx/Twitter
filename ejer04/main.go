package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var leyenda string
	var numero1 int
	var numero2 int
	fmt.Println("ingrese numero 1:")
	fmt.Scanf("%d", &numero1)

	fmt.Println("ingrese numero 2:")
	fmt.Scanf("%d", &numero2)
	fmt.Println(numero1, numero2)
	fmt.Println("ingrese descripcion:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	resultado := numero1 + numero2
	fmt.Println(leyenda, resultado)

}
