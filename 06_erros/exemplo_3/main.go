package main

import (
	"fmt"
)

func f() (string, string) {
	return "João 3:16", "Salmo 83:18"
}

func main() {
	_, v := f() // descartando a primeira variável de retorno
	fmt.Println(v)

	v, _ = f() // descartando a segunda variável de retorno
	fmt.Println(v)
}
