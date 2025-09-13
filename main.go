package main

// Every arquivo Go deve ter um pacote declarado na PRIMEIRA LINHA

import (
	"fmt"
	"myFirstGoProject/pacote"
)

func main() {
	fmt.Println("Hello World")
	fmt.Printf(pacote.Bar)
	pacote.PrintMinha()
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
