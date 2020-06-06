package main

import (
	"fmt"
	"strconv"
)

var numero int
var texto string
var status bool
var test byte

func main() {
	var numero2, numero3, numero1 int
	var moneda int64 = 0
	numero2, numero1, numero3, texto := 2, 10, 15, "esto es un texto"
	fmt.Println(numero1, numero2, numero3, texto)
	texto = strconv.Itoa(int(moneda))
	fmt.Println(texto)
}
