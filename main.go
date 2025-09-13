package main

// Every arquivo Go deve ter um pacote declarado na PRIMEIRA LINHA

import (
	"fmt"
	"myFirstGoProject/pacote"
	"myFirstGoProject/variaveis"
)

func main() {
	fmt.Println("Hello World")
	fmt.Printf("%s", pacote.Bar)
	pacote.PrintMinha()
	digaOi()
	fmt.Println(somar(2, 43))
	a, b := swap(10, 20)
	fmt.Println(a, b)

	res, rem := dividir(7, 2)
	fmt.Println(res, rem)

	x := somar2(2)(1)

	fmt.Println(x)

	variaveis.Foo()
}

func multiplicar(nums ...int) int {
	var out int
	for _, n := range nums {
		out *= n
	}

	return out
}

// HighOrder Functions

func somar2(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func dividir(a, b int) (int, int) {
	res := a / b
	rem := a % b
	return res, rem
}

func digaOi() {
	fmt.Println("Oi")
}

func somar(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

// var Foo string

// const Bar string = "bar"

// type Types struct {}

// Nome Publico e Privados

/** toda variavel, funcao, tipo, metodo, however, que iniciar com a letra MINUSCULA eh privado!
 */
func foo() {}

// Já toda a variavel que começar com a letra MAIUSCULA eh publico!
func Bar() {}
