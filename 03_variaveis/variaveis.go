package main

import "fmt"

func main() {

	/* Declarando e atribuindo variáveis */
	var a string = "Declaração longa"
	b := "Declaração curta"

	fmt.Printf("%v \n", a) /* %v imprime o valor e \n pula uma linha */
	fmt.Printf("%T \n", a) /* %T imprime o Tipo da variável */

	fmt.Printf("%v \n", b)
	fmt.Printf("%T \n", b)
}
