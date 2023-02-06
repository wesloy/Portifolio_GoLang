package main

import "fmt"

var num1 = 0

const name = "Números"

func numerosPrimos() {
	const name = "Números Primos"
	num1 := 3
	num2 := 7

	fmt.Println(name)
	fmt.Println(num1)
	fmt.Println(num2)
}

func main() {
	numerosPrimos()
	fmt.Println(name)
	fmt.Println(num1)

}
