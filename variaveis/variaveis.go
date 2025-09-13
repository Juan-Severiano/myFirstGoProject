package variaveis

import "fmt"

var idade int = 17

func Foo() {
	// var nome, sobrenome string = "Francisco Juan", "Severiano"

	// var (
	// 	name     string = "Camila"
	// 	lastName string = "Carneiro"
	// )

	// alias para var nome = ""
	nome := "Juan"
	sobrenome := "Severiano"

	fmt.Println(nome, sobrenome, idade)
}
